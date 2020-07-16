package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

// Error
var (
	ErrIDNotFound    = errors.New("id not found")
	ErrNodeIsRunning = errors.New("worker is running")
)

// worker status. the worker is in either idle, or running state
type status string

const (
	Idle    status = "idle"
	Running status = "running"
)

type unit struct {
	Title    string             // metric title
	Type     metrics.MetricType // to know the current unit type
	metricID int                // metric table foreign key
	c        gometrics.Counter
	h        gometrics.Histogram
	g        gometrics.Gauge
}

type metricLogger interface {
	Counter(context.Context, string, string, int64, int64) error
	Histogram(context.Context, string, string, int64, gometrics.Histogram) error
	Gauge(context.Context, string, string, int64, int64) error
	NewGroup(context.Context, metrics.Group) (*ent.Group, bool, error)
	NewGraph(context.Context, metrics.Graph, int) (*ent.Graph, error)
}

// Worker is the main structure for a running worker
// contains host information, the scenario (plugin)
// and gometrics unit
type Worker struct {
	mu       sync.Mutex
	id       string
	hostname string
	pid      int

	status     status
	pluginPath string
	vus        *scenario.Vus
	cancel     context.CancelFunc

	units map[string]unit // title - gometrics

	log metricLogger
}

// the singleton worker variable
var worker Worker

func init() {
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	// id return the identification of the worker which is the combination of
	// hostname and pid
	id := fmt.Sprintf("%s-%d", hostname, pid)

	worker = Worker{
		id:       id,
		pid:      pid,
		hostname: hostname,
		status:   Idle,

		units: make(map[string]unit),
	}
}

// NewWorker returns the singleton worker
func NewWorker(log metricLogger) (*Worker, error) {
	worker.log = log
	return &worker, nil
}

func (w *Worker) reset() {
	w.mu.Lock()
	w.status = Idle
	w.units = make(map[string]unit)
	w.cancel = nil
	w.mu.Unlock()
}

// Load downloads the go plugin, extracts the virtual user scenario
func (w *Worker) Load(so string) error {
	vus, err := scenario.LoadPlugin(so)
	if err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	w.pluginPath = so
	w.vus = &vus

	return nil
}

// Cancel stops the running scenario if any
func (w *Worker) Cancel() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.status == Idle {
		return nil
	}

	// if there is a running scenario
	w.cancel()

	return nil
}

// Run starts the preloaded plugin
// return error if the worker is running already
func (w *Worker) Run() error {
	w.mu.Lock()

	if w.status == Running {
		return ErrNodeIsRunning
	}

	w.status = Running
	w.mu.Unlock()

	w.run()
	return nil
}

func (w *Worker) run() {
	ctx, cancel := context.WithCancel(context.Background())

	w.cancel = cancel

	finished := make(chan struct{})

	go w.logScaled(ctx, 5*time.Second)
	go w.runScen(ctx, finished)

	select {
	case <-finished:
		log.Printf("scenarios finished")
	case <-ctx.Done():
		log.Printf("scenarios cancel")
	}

	// when finish, reset the worker
	w.reset()
}

func (w *Worker) Running() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.status == Running
}

func (w *Worker) runScen(ctx context.Context, done chan struct{}) {
	var totalVu int

	vus := *w.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	var wait sync.WaitGroup
	wait.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(i, j int) {
				vus[i].Fu(ctx, j)
				wait.Done()
			}(i, j)
		}
	}

	wait.Wait()

	done <- struct{}{}
}

// logScaled extract the metric log from a worker
// should run this function in a routine
func (w *Worker) logScaled(ctx context.Context, freq time.Duration) {
	ch := make(chan interface{})

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	if err := w.logScaledOnCue(ctx, ch); err != nil {
		log.Fatalln(err)
	}
}

func (w *Worker) logScaledOnCue(ctx context.Context, ch chan interface{}) error {
	for {
		select {
		case <-ch:
			now := timestampMs()
			w.mu.Lock()
			units := w.units
			w.mu.Unlock()

			for _, u := range units {
				switch u.Type {
				case metrics.Counter:
					w.log.Counter(ctx, w.id, u.Title, now, u.c.Count())
				case metrics.Histogram:
					h := u.h.Snapshot()
					w.log.Histogram(ctx, w.id, u.Title, now, h)
				case metrics.Gauge:
					w.log.Gauge(ctx, w.id, u.Title, now, u.g.Value())
				}
			}
		case <-ctx.Done():
			log.Printf("logScaledOnCue cancel")
			return nil
		}
	}

	return nil
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}

// Setup is used for the worker to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	ctx := context.TODO()

	units := make(map[string]unit)

	for _, group := range groups {
		// create a new group if not existed
		eg, created, err := worker.log.NewGroup(ctx, group)
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}
		// if the group is existed, continue
		if !created {
			continue
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			_, err := worker.log.NewGraph(ctx, graph, eg.ID)
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
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
						Title: m.Title,
						Type:  m.Type,
						// metricID: metricInstance.ID,
						c: c,
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
						Title: m.Title,
						Type:  m.Type,
						// metricID: metricInstance.ID,
						h: h,
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
						Title: m.Title,
						Type:  m.Type,
						// metricID: metricInstance.ID,
						g: g,
					}
				}
			}
		}
	}

	// aggregate units
	worker.mu.Lock()
	for k, v := range units {
		worker.units[k] = v
	}
	worker.mu.Unlock()

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	log.Printf("worker notify title: %s, value %d\n", title, value)

	worker.mu.Lock()
	defer worker.mu.Unlock()

	u, ok := worker.units[title]
	if !ok {
		log.Printf("error metric title %s not found\n", title)
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
