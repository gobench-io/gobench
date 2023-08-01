package executor

import (
	"context"

	api "github.com/gobench-io/gobench/v2/gen/go/pb"
	"google.golang.org/grpc"
)

// nop metric logger, used for testing
type nopLog struct{}

func (n *nopLog) Counter(ctx context.Context, req *api.CounterReq, opts ...grpc.CallOption) (*api.CounterRes, error) {
	return nil, nil
}

func (n *nopLog) Histogram(ctx context.Context, req *api.HistogramReq, opts ...grpc.CallOption) (*api.HistogramRes, error) {
	return nil, nil
}

func (n *nopLog) Gauge(ctx context.Context, req *api.GaugeReq, opts ...grpc.CallOption) (*api.GaugeRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateGroup(ctx context.Context, req *api.FCGroupReq, opts ...grpc.CallOption) (*api.FCGroupRes, error) {
	return new(api.FCGroupRes), nil
}

func (n *nopLog) FindCreateGraph(ctx context.Context, req *api.FCGraphReq, opts ...grpc.CallOption) (*api.FCGraphRes, error) {
	return new(api.FCGraphRes), nil
}

func (n *nopLog) FindCreateMetric(ctx context.Context, req *api.FCMetricReq, opts ...grpc.CallOption) (*api.FCMetricRes, error) {
	return new(api.FCMetricRes), nil
}

// NewNopMetricLog returns a no-op metric logger
func newNopMetricLog() *nopLog {
	return &nopLog{}
}
