package agent

import (
	"context"

	api "github.com/gobench-io/gobench/v2/gen/go/pb"
)

// nop metric logger
type nopLog struct{}

func (n *nopLog) FindCreateGroup(context.Context, *api.FCGroupReq) (*api.FCGroupRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateGraph(context.Context, *api.FCGraphReq) (*api.FCGraphRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateMetric(context.Context, *api.FCMetricReq) (*api.FCMetricRes, error) {
	return nil, nil
}

func (n *nopLog) Histogram(context.Context, *api.HistogramReq) (*api.HistogramRes, error) {
	return nil, nil
}

func (n *nopLog) Counter(context.Context, *api.CounterReq) (*api.CounterRes, error) {
	return nil, nil
}

func (n *nopLog) Gauge(context.Context, *api.GaugeReq) (*api.GaugeRes, error) {
	return nil, nil
}

func newNopMetricLog() *nopLog {
	return &nopLog{}
}
