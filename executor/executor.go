package executor

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/pb"
	"github.com/gobench-io/gobench/scenario"
	"google.golang.org/grpc"

	gometrics "github.com/rcrowley/go-metrics"
)

// executor status. The executor is in either idle, or running state
type status string

const (
	Idle    status = "idle"
	Running status = "running"
)

// Error
var (
	ErrIDNotFound    = errors.New("id not found")
	ErrNodeIsRunning = errors.New("driver is running")

	ErrAppCancel = errors.New("application is cancel")
	ErrAppPanic  = errors.New("application is panic")
)

type unit struct {
	Title    string             // metric title
	Type     metrics.MetricType // to know the current unit type
	metricID int                // metric table foreign key
	c        gometrics.Counter
	h        gometrics.Histogram
	g        gometrics.Gauge
}

// Options is for creating new executor object
type Options struct {
	AgentSock    string
	ExecutorSock string
	AppID        int
	Vus          scenario.Vus
}

// Executor struct
type Executor struct {
	mu           sync.Mutex
	id           string
	logger       logger.Logger
	agentSock    string
	executorSock string
	appID        int

	status status
	vus    scenario.Vus
	units  map[string]unit //title - gometrics

	rc pb.AgentClient
}

// the singleton instance of executor
var executorInstance Executor

func init() {
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	id := fmt.Sprintf("%s-%d", hostname, pid)

	executorInstance = Executor{
		id: id,
	}
}

func getExecutor() *Executor {
	return &executorInstance
}

// NewExecutor creates a new executor
// also load the plugin from driver path
func NewExecutor(opts *Options, logger logger.Logger) (e *Executor, err error) {
	e = getExecutor()

	e.units = make(map[string]unit)
	e.logger = logger
	e.agentSock = opts.AgentSock
	e.executorSock = opts.ExecutorSock
	e.appID = opts.AppID
	e.vus = opts.Vus

	e.status = Idle

	// establishes a connection to agent rpc server
	socket := "passthrough:///unix://" + e.agentSock
	conn, err := grpc.Dial(socket, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	e.rc = pb.NewAgentClient(conn)

	return
}

// Serve starts a rpc server at the executor socket
// and connects to the agent via agent socket
func (e *Executor) Serve() (err error) {
	// executor register a rpc server at executor socket
	l, err := net.Listen("unix", e.executorSock)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, e)

	err = s.Serve(l)
	if err != nil {
		return
	}

	return
}

func (e *Executor) run(ctx context.Context) (err error) {
	// first, setup executor's system load
	if err = e.systemloadSetup(); err != nil {
		return
	}

	// todo: check the status
	if e.status == Running {
		return ErrNodeIsRunning
	}

	e.status = Running

	finished := make(chan error)

	// when the runScen finished, we should stop the logScaled and systemloadRun
	// also; however, not necessary since the executor will be shutdown anyway
	go e.logScaled(ctx, 10*time.Second)
	go e.runScen(ctx, finished)
	go e.systemloadRun(ctx)

	select {
	case err = <-finished:
	case <-ctx.Done():
		err = ErrAppCancel
	}

	// todo: update status

	return
}

func (e *Executor) runScen(ctx context.Context, done chan<- error) {
	var totalVu int

	vus := e.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	var wg sync.WaitGroup
	wg.Add(totalVu)

	for i := range vus {
		go func(i int) {
			for j := 0; j < vus[i].Nu; j++ {
				go func(i, j int) {
					vus[i].Fu(ctx, j)
					wg.Done()
				}(i, j)
				dis.SleepRatePoisson(vus[i].Rate)
			}
		}(i)
	}

	wg.Wait()
	done <- nil
}

// logScaled extract the metric log from a driver
// should run this function in a routine
func (e *Executor) logScaled(ctx context.Context, freq time.Duration) {
	ch := make(chan interface{})

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	if err := e.logScaledOnCue(ctx, ch); err != nil {
		e.logger.Fatalw("failed logScaledOnCue", "err", err)
	}
}

func (e *Executor) logScaledOnCue(ctx context.Context, ch chan interface{}) error {
	var err error
	for {
		select {
		case <-ch:
			now := timestampMs()
			e.mu.Lock()
			units := e.units
			e.mu.Unlock()

			for _, u := range units {
				base := &pb.BasedReqMetric{
					AppID: int64(e.appID),
					EID:   e.id,
					MID:   int64(u.metricID),
					Time:  now,
				}

				switch u.Type {
				case metrics.Counter:
					_, err = e.rc.Counter(ctx, &pb.CounterReq{
						Base:  base,
						Count: u.c.Count(),
					})
				case metrics.Histogram:
					h := u.h.Snapshot()
					ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
					hv := &pb.HistogramValues{
						Count:  h.Count(),
						Min:    h.Min(),
						Max:    h.Max(),
						Mean:   h.Mean(),
						Stddev: h.StdDev(),
						Median: ps[0],
						P75:    ps[1],
						P95:    ps[2],
						P99:    ps[3],
						P999:   ps[4],
					}
					_, err = e.rc.Histogram(ctx, &pb.HistogramReq{
						Base:      base,
						Histogram: hv,
					})
				case metrics.Gauge:
					_, err = e.rc.Gauge(ctx, &pb.GaugeReq{
						Base:  base,
						Gauge: u.g.Value(),
					})
				}

				if err != nil {
					e.logger.Errorw("metric log failed", "err", err)
				}
			}
		case <-ctx.Done():
			e.logger.Infow("logScaledOnCue canceled")
			return nil
		}
	}
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}
