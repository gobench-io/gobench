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
	vus *scenario.Vus
}

const (
	idle    status = "idle"
	running status = "running"
)

func New() (*Node, err) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	id := os.Getpid()

	id = fmt.Sprintf("%s-%i", hostname, pid)

	return &Node{
		id: id,
		status: idle,
	}
}

func (n *Node) Load(so string) err {
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
