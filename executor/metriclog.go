package executor

import (
	"context"
	"fmt"

	"github.com/gobench-io/gobench/pb"
)

// metricLog interface implementer for the Executor

func (e *Executor) Counter(ctx context.Context, mID int, title string, time, c int64) (
	err error,
) {
	_, err = e.rc.Counter(ctx, &pb.CounterReq{
		Base: &pb.BasedReqMetric{
			AppID: int64(e.appID),
			EID:   e.id,
			MID:   int64(mID),
			Time:  time,
		},
		Count: c,
	})
	if err != nil {
		err = fmt.Errorf("rpc counter: %v", err)
		return
	}

	return
}

func (e *Executor) Histogram(ctx context.Context, mID int, title string, time int64, h *pb.HistogramValues) (
	err error,
) {
	_, err = e.rc.Histogram(ctx, &pb.HistogramReq{
		Base: &pb.BasedReqMetric{
			AppID: int64(e.appID),
			EID:   e.id,
			MID:   int64(mID),
			Time:  time,
		},
		Histogram: h,
	})

	if err != nil {
		err = fmt.Errorf("rpc histogram: %v", err)
		return
	}

	return
}

func (e *Executor) Gauge(ctx context.Context, mID int, title string, time int64, g int64) (
	err error,
) {
	_, err = e.rc.Gauge(ctx, &pb.GaugeReq{
		Base: &pb.BasedReqMetric{
			AppID: int64(e.appID),
			EID:   e.id,
			MID:   int64(mID),
			Time:  time,
		},
		Gauge: g,
	})
	if err != nil {
		err = fmt.Errorf("rpc gauge: %v", err)
		return
	}

	return
}

func (e *Executor) FindCreateGroup(ctx context.Context, req *pb.FCGroupReq) (
	res *pb.FCGroupRes, err error,
) {
	return e.rc.FindCreateGroup(ctx, req)
}

func (e *Executor) FindCreateGraph(ctx context.Context, req *pb.FCGraphReq) (res *pb.FCGraphRes, err error) {
	return e.rc.FindCreateGraph(ctx, req)
}

func (e *Executor) FindCreateMetric(ctx context.Context, req *pb.FCMetricReq) (*pb.FCMetricRes, error) {
	return e.rc.FindCreateMetric(ctx, req)
}
