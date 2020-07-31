package worker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

// Error
var (
	ErrIDNotFound    = errors.New("id not found")
	ErrNodeIsRunning = errors.New("worker is running")

	ErrAppCancel = errors.New("application is cancel")
	ErrAppPanic  = errors.New("application is panic")
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
	Counter(context.Context, int, string, string, int64, int64) error
	Histogram(context.Context, int, string, string, int64, gometrics.Histogram) error
	Gauge(context.Context, int, string, string, int64, int64) error
	FindCreateGroup(context.Context, metrics.Group, int) (*ent.Group, error)
	FindCreateGraph(context.Context, metrics.Graph, int) (*ent.Graph, error)
	FindCreateMetric(context.Context, metrics.Metric, int) (*ent.Metric, error)
}

// Worker is the main structure for a running worker
// contains host information, the scenario (plugin)
// and gometrics unit
type Worker struct {
	mu       sync.Mutex
	id       string
	hostname string
	pid      int

	appID      int
	status     status
	pluginPath string
	vus        *scenario.Vus

	units map[string]unit // title - gometrics

	logger logger.Logger
	ml     metricLogger
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
	}
}

// NewWorker returns the singleton worker
func NewWorker(ml metricLogger, logger logger.Logger, appID int) (*Worker, error) {
	worker.ml = ml
	worker.logger = logger
	worker.units = make(map[string]unit)
	worker.appID = appID

	// reset metrics
	worker.unregisterGometrics()

	return &worker, nil
}

func (w *Worker) unregisterGometrics() {
	gometrics.Each(func(name string, i interface{}) {
		gometrics.Unregister(name)
	})
}

func (w *Worker) reset() {
	w.mu.Lock()
	w.status = Idle
	w.units = make(map[string]unit)
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

// Run starts the preloaded plugin
// return error if the worker is running already
func (w *Worker) Run(ctx context.Context) (err error) {
	w.mu.Lock()

	if w.status == Running {
		w.mu.Unlock()
		return ErrNodeIsRunning
	}

	w.status = Running
	w.mu.Unlock()

	err = w.run(ctx)

	return err
}

func (w *Worker) run(ctx context.Context) (err error) {
	finished := make(chan error)

	go w.logScaled(ctx, 10*time.Second)
	go w.runScen(ctx, finished)

	select {
	case err = <-finished:
	case <-ctx.Done():
		err = ErrAppCancel
	}

	// when finish, reset the worker
	w.reset()

	return
}

// Running returns a bool value indicating that the working is running
func (w *Worker) Running() bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.status == Running
}

func (w *Worker) runScen(ctx context.Context, done chan error) {
	var totalVu int

	vus := *w.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	fatalErr := make(chan error)
	wgDone := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(i, j int) {

				defer func() {
					if r := recover(); r != nil {
						w.logger.Errorw("recovered in runScreen",
							"err", r,
						)
						fatalErr <- ErrAppPanic
						// return
					}
				}()

				vus[i].Fu(ctx, j)
				wg.Done()
			}(i, j)
		}
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		done <- nil
	case err := <-fatalErr:
		done <- err
	}
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
		w.logger.Fatalw("failed logScaledOnCue", "err", err)
	}
}

func (w *Worker) logScaledOnCue(ctx context.Context, ch chan interface{}) error {
	var err error
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
					err = w.ml.Counter(ctx, u.metricID, w.id, u.Title, now, u.c.Count())
				case metrics.Histogram:
					h := u.h.Snapshot()
					err = w.ml.Histogram(ctx, u.metricID, w.id, u.Title, now, h)
				case metrics.Gauge:
					err = w.ml.Gauge(ctx, u.metricID, w.id, u.Title, now, u.g.Value())
				}
				if err != nil {
					w.logger.Errorw("metric log failed",
						"err", err,
					)
				}
			}
		case <-ctx.Done():
			w.logger.Infow("logScaledOnCue canceled")
			return nil
		}
	}
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}

// Setup is used for the worker to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	ctx := context.TODO()

	units := make(map[string]unit)

	worker.mu.Lock()
	defer worker.mu.Unlock()

	for _, group := range groups {
		// create a new group if not existed
		egroup, err := worker.ml.FindCreateGroup(ctx, group, worker.appID)
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			egraph, err := worker.ml.FindCreateGraph(ctx, graph, egroup.ID)
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				emetric, err := worker.ml.FindCreateMetric(ctx, m, egraph.ID)
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
						metricID: emetric.ID,
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
						metricID: emetric.ID,
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
						metricID: emetric.ID,
						g:        g,
					}
				}
			}
		}
	}

	// aggregate units
	for k, v := range units {
		worker.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	worker.mu.Lock()
	defer worker.mu.Unlock()

	u, ok := worker.units[title]
	if !ok {
		worker.logger.Infow("metric not found",
			"title", title,
		)
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
