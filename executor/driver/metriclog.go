package driver

import (
	"context"

	"github.com/gobench-io/gobench/pb"
	"google.golang.org/grpc"
)

// nil metric logger
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

func (n *nopLog) FindCreateGroup(ctx context.Context, req *pb.FCGroupReq, opts ...grpc.CallOption)(*pb.FCGroupRes, error) {
	return new(pb.FCGroupRes), nil
}

func (n *nopLog) FindCreateGraph(context.Context, *pb.FCGraphReq) (*, opts ...grpc.CallOptionpb.FCGraphRes, error) {
	return new(pb.FCGraphRes), nil
}
func (n *nopLog) FindCreateMetric(context.Context, *pb.FCMetricReq) (*, opts ...grpc.CallOptionpb.FCMetricRes, error) {
	return new(pb.FCMetricRes), nil
}

// NewNopMetricLog returns a no-op metric logger
func newNopMetricLog() *nopLog {
	return &nopLog{}
}
