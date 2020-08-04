package agent

import "github.com/gobench-io/gobench/metrics"

type metricLoggerRPC interface {
	FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error
	FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error
	FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error
}

type Agent struct {
	ml metricLoggerRPC
}

func NewAgent(ml metricLoggerRPC) (*Agent, error) {
	a := *&Agent{
		ml: ml,
	}
	return &a, nil
}
