package node

import (
	"os"
	"sync"

	"github.com/gobench-io/gobench/scenario"
)

type status string

type Node struct {
	mu         sync.Mutex
	id         string
	status     status
	pluginPath string
	vus        *scenario.Vus
}

const (
	idle    status = "idle"
	running status = "running"
)

func New() (*Node, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	id := os.Getpid()

	id = fmt.Sprintf("%s-%i", hostname, pid)

	return &Node{
		id:     id,
		status: idle,
	}
}

func (n *Node) Load(so string) error {
	vus, err := scenario.LoadPlugin(so)
	if err != nil {
		return err
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.pluginPath = so
	n.vus = &vus
}

func (n *Node) Run() {
	n.mu.Lock()
	n.status = running
	n.mu.Unlock()
}

func (n *Node) run() {
	var donewg sync.WaitGroup

	var totalVu int

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

	n.mu.Lock()
	n.status = idle
	n.mu.Unlock()
}
