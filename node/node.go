package node

import (
	"fmt"
	"os"
	"sync"

	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

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
	units      map[string]unit
}

const (
	idle    status = "idle"
	running status = "running"
)

// the sigleton node variable
var node Node

func init() {
	hostname, _:= os.Hostname()
	pid := os.Getpid()

	node = Node{
		pid:      pid,
		hostname: hostname,
		status:   idle,
	}
}

// New return the singleton node
func New() (*Node, error) {
	return &node, nil
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
	var donewg sync.WaitGroup

	var totalVu int

	vus := *n.vus

	for i := range vus {
		totalVu += vus[i].Nu
	}

	donewg.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(j int) {
				vus[i].Fu(j, &donewg)
			}(j)
		}
	}

	donewg.Wait()

	// when finish, set node to idle
	n.mu.Lock()
	n.status = idle
	n.mu.Unlock()
}
