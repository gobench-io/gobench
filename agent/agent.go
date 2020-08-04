package agent

import "github.com/gobench-io/gobench/metrics"

type metricLoggerRPC interface {
	FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error
	FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error
	FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error
}

type Agent struct {
	// when local agent, inherit from master
	// when remote agent, ... todo
	metricLoggerRPC
}

func NewAgent(ml metricLoggerRPC) (*Agent, error) {
	a := *&Agent{
		metricLoggerRPC: ml,
	}
	return &a, nil
}
