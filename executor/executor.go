package executor

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/gobench-io/gobench/v2/dis"
	"github.com/gobench-io/gobench/v2/executor/metrics"
	"github.com/gobench-io/gobench/v2/executor/scenario"
	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/pb"
	"google.golang.org/grpc"

	gometrics "github.com/rcrowley/go-metrics"
)

// executor status. The executor is in either idle, or running state
type status string

const (
	Idle     status = "idle"
	Running  status = "running"
	Finished status = "finished"
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

	return
}

// Serve starts a rpc server at the executor socket
// and connects to the agent via agent socket
func (e *Executor) Serve() (err error) {
	// establishes a connection to agent rpc server
	socket := "passthrough:///unix://" + e.agentSock
	conn, err := grpc.Dial(socket, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	e.rc = pb.NewAgentClient(conn)

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
	e.status = Finished

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

// Setup is used for the driver to report the metrics that it will generate
func (e *Executor) Setup(groups []metrics.Group) error {
	ctx := context.TODO()

	units := make(map[string]unit)

	e.mu.Lock()
	defer e.mu.Unlock()

	for _, group := range groups {
		// create a new group if not existed
		egroup, err := e.rc.FindCreateGroup(ctx, &pb.FCGroupReq{
			AppID: int64(e.appID),
			Name:  group.Name,
		})
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			egraph, err := e.rc.FindCreateGraph(ctx, &pb.FCGraphReq{
				AppID:   int64(e.appID),
				Title:   graph.Title,
				Unit:    graph.Unit,
				GroupID: egroup.Id,
			})
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				emetric, err := e.rc.FindCreateMetric(ctx, &pb.FCMetricReq{
					AppID:   int64(e.appID),
					Title:   m.Title,
					Type:    string(m.Type),
					GraphID: egraph.Id,
				})
				if err != nil {
					return fmt.Errorf("failed create metric: %v", err)
				}

				// counter type
				if m.Type == metrics.Counter {
					c := gometrics.NewCounter()
					if err := gometrics.Register(m.Title, c); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}

					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						c:        c,
					}
				}

				if m.Type == metrics.Histogram {
					s := gometrics.NewExpDecaySample(1028, 0.015)
					h := gometrics.NewHistogram(s)
					if err := gometrics.Register(m.Title, h); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						h:        h,
					}
				}

				if m.Type == metrics.Gauge {
					g := gometrics.NewGauge()
					if err := gometrics.Register(m.Title, g); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						g:        g,
					}
				}
			}
		}
	}

	// aggregate units
	for k, v := range units {
		e.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func (e *Executor) Notify(title string, value int64) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	u, ok := e.units[title]
	if !ok {
		e.logger.Infow("metric not found", "title", title)
		return ErrIDNotFound
	}

	if u.Type == metrics.Counter {
		u.c.Inc(value)
	}

	if u.Type == metrics.Histogram {
		u.h.Update(value)
	}

	if u.Type == metrics.Gauge {
		u.g.Update(value)
	}

	return nil
}
