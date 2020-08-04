package executor

import (
	"context"
	"fmt"

	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"
)

// metricLog interface implementer for the Executor

func (e *Executor) Counter(ctx context.Context, mID int, title string, time, c int64) error {
	return nil
}

func (e *Executor) Histogram(ctx context.Context, mID int, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (e *Executor) Gauge(ctx context.Context, mID int, title string, time int64, g int64) error {
	return nil
}

func (e *Executor) FindCreateGroup(ctx context.Context, mg metrics.Group, appID int) (
	res *metrics.FCGroupRes, err error,
) {
	res = new(metrics.FCGroupRes)

	req := &metrics.FCGroupReq{
		Name:  mg.Name,
		AppID: appID,
	}
	// todo: should call Agent.FindCreateGroupRPC
	if err = e.rc.Call("Master.FindCreateGroupRPC", req, res); err != nil {
		err = fmt.Errorf("rpc FindCreateGroup: %v", err)
		return
	}
	return
}

func (e *Executor) FindCreateGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	res *metrics.FCGraphRes, err error,
) {
	return nil, nil
}

func (e *Executor) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	res *metrics.FCMetricRes, err error,
) {
	return nil, nil
}
