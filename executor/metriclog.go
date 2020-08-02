package main

import (
	"context"

	"github.com/gobench-io/gobench/ent"
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
	eg *ent.Group, err error,
) {
	return nil, nil
}

func (e *Executor) FindCreateGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	egraph *ent.Graph, err error,
) {
	return nil, nil
}

func (e *Executor) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	emetric *ent.Metric, err error,
) {
	return nil, nil
}
