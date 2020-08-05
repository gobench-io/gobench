package agent

import (
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
	rs     *rpc.Server
}

func NewAgent(ml metricLoggerRPC) (*Agent, error) {
	a := &Agent{
		metricLoggerRPC: ml,
		rs:              rpc.NewServer(),
	}
	return a, nil
}

func (a *Agent) StartSocketServer(socket string) error {
	a.socket = socket

	a.rs.Register(a)

	serverMux := http.NewServeMux()
	serverMux.Handle(rpc.DefaultRPCPath, a.rs)
	serverMux.Handle(rpc.DefaultDebugPath, a.rs)

	os.Remove(socket)
	l, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}

	go http.Serve(l, serverMux)

	return nil
}

func (a *Agent) GetSocketName() string {
	return a.socket
}
