package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"

	"github.com/gobench-io/gobench/services/smtp"
)

var ErrIdNotFound = errors.New("id not found")

type BenchStatus string

const (
	StatusInit     BenchStatus = "init"
	StatusRunning  BenchStatus = "running"
	StatusFinished BenchStatus = "finished"
	StatusCancel   BenchStatus = "cancel"
)

type unit struct {
	Title    string             // metric title
	Type     metrics.MetricType // to know the currnt unit type
	metricID int                // metric table foreign key
	c        gometrics.Counter
	h        gometrics.Histogram
	g        gometrics.Gauge
}

type Collect struct {
	mu         sync.Mutex
	units      map[string]unit
	smtp       smtp.Service
	dbFilename string

	name       string
	status     BenchStatus
	createdAt  time.Time
	finishedAt time.Time

	DB *ent.Client
}

var benchCollect Collect

func init() {
	benchCollect = Collect{
		units:     make(map[string]unit),
		status:    StatusInit,
		createdAt: time.Now(),
	}
}

// NewBench returns the singleton instance of the benchmark
func NewBench() *Collect {
	return &benchCollect
}

// GetInstance returns the singleton instance of the benchmark
func GetInstance() *Collect {
	return &benchCollect
}

// Setup starts the gobench program
// by creating a collect entity and database to store units
func Setup(groups []metrics.Group) error {
	units := make(map[string]unit)
	ctx := context.Background()

	benchCollect.mu.Lock()
	defer benchCollect.mu.Unlock()

	for _, group := range groups {
		// create new group if not existed
		groupInst, created, err := bFindOrCreateGroup(ctx, benchCollect.DB, group)
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}
		// if the group is existed, continue
		if !created {
			continue
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			graphInst, err := bCreateGraph(ctx, benchCollect.DB, graph, groupInst)
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				metricInstance, err := bFindOrCreateMetric(ctx, benchCollect.DB, m, graphInst)
				if err != nil {
					return fmt.Errorf("failed create metric: %v", err)
				}

				// create unit map
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
						metricID: metricInstance.ID,
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
						metricID: metricInstance.ID,
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
						metricID: metricInstance.ID,
						g:        g,
					}
				}
			}
		}
	}

	// aggregrate units
	for k, v := range units {
		benchCollect.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
func Notify(title string, value int64) error {
	u, ok := benchCollect.units[title]
	if !ok {
		log.Printf("error metric title %s not found\n", title)
		return ErrIdNotFound
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

// Start function setup the db with benchCollect
// and begin to log metrics to db
func (c *Collect) Start() error {
	log.Println("benchmark starting")

	c.status = StatusRunning

	filename := time.Now().Format("2006-01-02-15-04-05") + ".sqlite3"
	if err := c.setupDb(filename); err != nil {
		return err
	}

	if err := c.createApplication(context.Background()); err != nil {
		return err
	}

	c.handleSignals()

	// log the metric in the background
	go c.logScaled(5 * time.Second)

	return nil
}

func (c *Collect) finish(s BenchStatus) error {
	log.Println("benchmark is shutting down")

	c.status = s
	c.finishedAt = time.Now()

	if err := c.updateApplication(context.Background()); err != nil {
		log.Printf("failed to update application %v\n", err)
	}

	if c.smtp.IsEnabled() {
		if err := c.smtp.Send(smtp.Result{
			Name:       c.name,
			CreatedAt:  c.createdAt,
			FinishedAt: c.finishedAt,
			Status:     string(c.status),
			FilePath:   c.dbFilename,
		}); err != nil {
			log.Printf("failed to send email %v\n", err)
		}
	}

	return c.DB.Close()
}

// Finish closes the db connection
// and trigger services
func (c *Collect) Finish() error {
	return c.finish(StatusFinished)
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}

func (c *Collect) logScaled(freq time.Duration) {
	ch := make(chan interface{})
	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)
	if err := c.logScaledOnCue(ch); err != nil {
		log.Fatalln(err)
	}
}

func (c *Collect) logScaledOnCue(ch chan interface{}) error {
	units := c.units
	ctx := context.Background()

	for range ch {
		now := timestampMs()
		for _, u := range units {
			switch u.Type {
			case metrics.Counter:
				_, err := c.DB.Counter.Create().
					SetCount(u.c.Count()).
					SetTime(now).
					SetMetricID(u.metricID).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed create count: %v", err)
				}

			case metrics.Histogram:
				h := u.h.Snapshot()
				ps := u.h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
				_, err := c.DB.Histogram.Create().
					SetCount(h.Count()).
					SetMin(h.Min()).
					SetMax(h.Max()).
					SetMean(h.Mean()).
					SetStddev(h.StdDev()).
					SetMedian(ps[0]).
					SetP75(ps[1]).
					SetP95(ps[2]).
					SetP99(ps[3]).
					SetP999(ps[4]).
					SetTime(now).
					SetMetricID(u.metricID).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed creating count: %v", err)
				}
			case metrics.Gauge:
				_, err := c.DB.Gauge.Create().
					SetValue(u.g.Value()).
					SetTime(now).
					SetMetricID(u.metricID).
					Save(ctx)
				if err != nil {
					return fmt.Errorf("failed creating gauge: %v", err)
				}
			}
		}
	}
	return nil
}

// Email sets the smtp service configuration
func (c *Collect) Email(config smtp.Config) *Collect {
	c.smtp = smtp.NewService(config)
	return c
}

func (c *Collect) Name(name string) *Collect {
	c.name = name
	return c
}
