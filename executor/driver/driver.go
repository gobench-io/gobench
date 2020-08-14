package driver

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

// Error
var (
	ErrIDNotFound    = errors.New("id not found")
	ErrNodeIsRunning = errors.New("driver is running")

	ErrAppCancel = errors.New("application is cancel")
	ErrAppPanic  = errors.New("application is panic")
)

// driver status. the driver is in either idle, or running state
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

// Driver is the main structure for a running driver
// contains host information, the scenario (plugin)
// and gometrics unit
type Driver struct {
	mu         sync.Mutex
	appID      int
	status     status
	driverPath string
	vus        *scenario.Vus

	units map[string]unit // title - gometrics

	logger logger.Logger
	ml     metricLogger
}

// the singleton driver variable
var driver Driver

func init() {
	driver = Driver{
		status: Idle,
	}
}

// NewDriver returns the singleton driver
func NewDriver(ml metricLogger, logger logger.Logger, driverPath string, appID int) (*Driver, error) {
	driver.mu.Lock()

	driver.ml = ml
	driver.logger = logger
	driver.units = make(map[string]unit)
	driver.appID = appID
	// reset metrics
	driver.unregisterGometrics()

	driver.mu.Unlock()

	err := driver.load(driverPath)

	return &driver, err
}

func (d *Driver) unregisterGometrics() {
	gometrics.Each(func(name string, i interface{}) {
		gometrics.Unregister(name)
	})
}

// load downloads the go plugin, extracts the virtual user scenario
func (d *Driver) load(so string) (err error) {
	vus, err := scenario.LoadPlugin(so)
	if err != nil {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.driverPath = so
	d.vus = &vus

	return
}

func (d *Driver) reset() {
	d.mu.Lock()
	d.status = Idle
	d.units = make(map[string]unit)
	d.mu.Unlock()
}

// SetNopMetricLog update the driver metric logger to the nop one. Mostly use
// for testing
func (d *Driver) SetNopMetricLog() error {
	nop := newNopMetricLog()

	d.mu.Lock()
	d.ml = nop
	d.mu.Unlock()

	return nil
}

// Run starts the preloaded plugin
// return error if the driver is running already
func (d *Driver) Run(ctx context.Context) (err error) {
	// first, setup driver system load
	if err = d.systemloadSetup(); err != nil {
		return
	}

	d.mu.Lock()

	if d.status == Running {
		d.mu.Unlock()
		return ErrNodeIsRunning
	}

	d.status = Running
	d.mu.Unlock()

	err = d.run(ctx)

	return err
}

func (d *Driver) run(ctx context.Context) (err error) {
	finished := make(chan error)

	// when the runScen finished, we should stop the logScaled and systemloadRun
	// also; however, not necessary since the executor will be shutdown anyway
	go d.logScaled(ctx, 10*time.Second)
	go d.runScen(ctx, finished)
	go d.systemloadRun(ctx)

	select {
	case err = <-finished:
	case <-ctx.Done():
		err = ErrAppCancel
	}

	// when finish, reset the driver
	d.reset()

	return
}

// Running returns a bool value indicating that the working is running
func (d *Driver) Running() bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.status == Running
}

func (d *Driver) runScen(ctx context.Context, done chan<- error) {
	var totalVu int

	vus := *d.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	var wg sync.WaitGroup
	wg.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(i, j int) {
				vus[i].Fu(ctx, j)
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait()
	done <- nil
}

// logScaled extract the metric log from a driver
// should run this function in a routine
func (d *Driver) logScaled(ctx context.Context, freq time.Duration) {
	ch := make(chan interface{})

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	if err := d.logScaledOnCue(ctx, ch); err != nil {
		d.logger.Fatalw("failed logScaledOnCue", "err", err)
	}
}

func (d *Driver) logScaledOnCue(ctx context.Context, ch chan interface{}) error {
	var err error
	for {
		select {
		case <-ch:
			now := timestampMs()
			d.mu.Lock()
			units := d.units
			d.mu.Unlock()

			for _, u := range units {
				switch u.Type {
				case metrics.Counter:
					err = d.ml.Counter(ctx, u.metricID, u.Title, now, u.c.Count())
				case metrics.Histogram:
					h := u.h.Snapshot()
					ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
					hv := metrics.HistogramValues{
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
					err = d.ml.Histogram(ctx, u.metricID, u.Title, now, hv)
				case metrics.Gauge:
					err = d.ml.Gauge(ctx, u.metricID, u.Title, now, u.g.Value())
				}

				if err != nil {
					d.logger.Errorw("metric log failed", "err", err)
				}
			}
		case <-ctx.Done():
			d.logger.Infow("logScaledOnCue canceled")
			return nil
		}
	}
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}

// Setup is used for the driver to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	ctx := context.TODO()

	units := make(map[string]unit)

	driver.mu.Lock()
	defer driver.mu.Unlock()

	for _, group := range groups {
		// create a new group if not existed
		egroup, err := driver.ml.FindCreateGroup(ctx, group, driver.appID)
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			egraph, err := driver.ml.FindCreateGraph(ctx, graph, egroup.ID)
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				emetric, err := driver.ml.FindCreateMetric(ctx, m, egraph.ID)
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
		driver.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	driver.mu.Lock()
	defer driver.mu.Unlock()

	u, ok := driver.units[title]
	if !ok {
		driver.logger.Infow("metric not found", "title", title)
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
