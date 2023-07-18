package executor

import (
	"context"

	"github.com/gobench-io/gobench/v2/pb"
	"google.golang.org/grpc"
)

// nop metric logger, used for testing
type nopLog struct{}

func (n *nopLog) Counter(ctx context.Context, req *pb.CounterReq, opts ...grpc.CallOption) (*pb.CounterRes, error) {
	return nil, nil
}

func (n *nopLog) Histogram(ctx context.Context, req *pb.HistogramReq, opts ...grpc.CallOption) (*pb.HistogramRes, error) {
	return nil, nil
}

func (n *nopLog) Gauge(ctx context.Context, req *pb.GaugeReq, opts ...grpc.CallOption) (*pb.GaugeRes, error) {
	return nil, nil
}

func (n *nopLog) FindCreateGroup(ctx context.Context, req *pb.FCGroupReq, opts ...grpc.CallOption) (*pb.FCGroupRes, error) {
	return new(pb.FCGroupRes), nil
}

func (n *nopLog) FindCreateGraph(ctx context.Context, req *pb.FCGraphReq, opts ...grpc.CallOption) (*pb.FCGraphRes, error) {
	return new(pb.FCGraphRes), nil
}

func (n *nopLog) FindCreateMetric(ctx context.Context, req *pb.FCMetricReq, opts ...grpc.CallOption) (*pb.FCMetricRes, error) {
	return new(pb.FCMetricRes), nil
}

// NewNopMetricLog returns a no-op metric logger
func newNopMetricLog() *nopLog {
	return &nopLog{}
}
