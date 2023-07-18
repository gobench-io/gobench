package agent

import (
	"context"

	"github.com/gobench-io/gobench/v2/pb"
)

// nop metric logger
type nopLog struct{}

func (n *nopLog) FindCreateGroup(context.Context, *pb.FCGroupReq) (*pb.FCGroupRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateGraph(context.Context, *pb.FCGraphReq) (*pb.FCGraphRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateMetric(context.Context, *pb.FCMetricReq) (*pb.FCMetricRes, error) {
	return nil, nil
}

func (n *nopLog) Histogram(context.Context, *pb.HistogramReq) (*pb.HistogramRes, error) {
	return nil, nil
}

func (n *nopLog) Counter(context.Context, *pb.CounterReq) (*pb.CounterRes, error) {
	return nil, nil
}

func (n *nopLog) Gauge(context.Context, *pb.GaugeReq) (*pb.GaugeRes, error) {
	return nil, nil
}

func newNopMetricLog() *nopLog {
	return &nopLog{}
}
