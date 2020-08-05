package agent

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/gobench-io/gobench/metrics"
)

type metricLoggerRPC interface {
	FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error
	FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error
	FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error
}

type Agent struct {
	// when this is the local agent, inherit from master
	// when this is the remote agent, ... todo
	metricLoggerRPC
	socket string
}

func NewAgent(ml metricLoggerRPC) (*Agent, error) {
	a := *&Agent{
		metricLoggerRPC: ml,
	}
	return &a, nil
}

func (a *Agent) StartSocketServer(socket string) error {
	log.Printf("-- default server: %p\n", rpc.DefaultServer)
	rpc.Register(a)
	rpc.HandleHTTP()
	os.Remove(socket)

	l, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}

	a.socket = socket

	go http.Serve(l, nil)

	return nil
}

func (a *Agent) GetSocketName() string {
	return a.socket
}
