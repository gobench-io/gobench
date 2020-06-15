package node

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

var ErrIdNotFound = errors.New("id not found")

type status string

type unit struct {
	Title    string             // metric title
	Type     metrics.MetricType // to know the currnt unit type
	metricID int                // metric table foreign key
	c        gometrics.Counter
	h        gometrics.Histogram
	g        gometrics.Gauge
}

type Node struct {
	mu       sync.Mutex
	hostname string
	pid      int

	status     status
	pluginPath string
	vus        *scenario.Vus

	units map[string]unit
}

const (
	idle    status = "idle"
	running status = "running"
)

// the sigleton node variable
var node Node

func init() {
	hostname, _ := os.Hostname()
	pid := os.Getpid()

	node = Node{
		pid:      pid,
		hostname: hostname,
		status:   idle,

		units: make(map[string]unit),
	}
}

// New return the singleton node
func New() (*Node, error) {
	return &node, nil
}

func (n *Node) reset() {
	n.mu.Lock()
	n.status = idle
	n.units = make(map[string]unit)
	n.mu.Unlock()
}

// Load downloads the go plugin, extracts the virtual user scenario
func (n *Node) Load(so string) error {
	vus, err := scenario.LoadPlugin(so)
	if err != nil {
		return err
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.pluginPath = so
	n.vus = &vus

	return nil
}

// Run starts the preloaded plugin
func (n *Node) Run() {
	n.mu.Lock()
	n.status = running
	n.mu.Unlock()

	n.run()
}

func (n *Node) run() {
	ctx, cancel := context.WithCancel(context.Background())

	var donewg sync.WaitGroup

	var totalVu int

	vus := *n.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	go n.logScaled(10 * time.Second)

	donewg.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(j int) {
				vus[i].Fu(ctx, j, &donewg)
			}(j)
		}
	}

	donewg.Wait()

	// when finish, reset the node
	n.reset()
}

func (n *Node) logScaled(freq time.Duration) {
	ch := make(chan interface{})

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	if err := n.logScaledOnCue(ch); err != nil {
		log.Fatalln(err)
	}
}

func (n *Node) logScaledOnCue(ch chan interface{}) error {
	log.Println("logScaledOnCue")
	return nil
}

// Setup is used for the worker to report the metrics that it will generate
func (n *Node) Setup(groups []metrics.Group) error {
	units := make(map[string]unit)

	for _, group := range groups {
		for _, graph := range group.Graphs {
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

	// aggregrate units
	for k, v := range units {
		n.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
func Notify(title string, value int64) error {
	node.mu.Lock()
	defer node.mu.Unlock()

	u, ok := node.units[title]
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
