package driver

import (
	"context"

	"github.com/gobench-io/gobench/metrics"
)

type metricLogger interface {
	Counter(context.Context, int, string, int64, int64) error
	Histogram(context.Context, int, string, int64, metrics.HistogramValues) error
	Gauge(context.Context, int, string, int64, int64) error
	FindCreateGroup(context.Context, metrics.Group, int) (*metrics.FCGroupRes, error)
	FindCreateGraph(context.Context, metrics.Graph, int) (*metrics.FCGraphRes, error)
	FindCreateMetric(context.Context, metrics.Metric, int) (*metrics.FCMetricRes, error)
}

// nil metric logger
type nilLog struct{}

func (l *nilLog) Counter(ctx context.Context, mID int, title string, time, c int64) error {
	return nil
}

func (l *nilLog) Histogram(ctx context.Context, mID int, title string, time int64, h metrics.HistogramValues) error {
	return nil
}

func (l *nilLog) Gauge(ctx context.Context, mID int, title string, time int64, g int64) error {
	return nil
}

func (l *nilLog) FindCreateGroup(ctx context.Context, mg metrics.Group, appID int) (
	*metrics.FCGroupRes, error,
) {

	return new(metrics.FCGroupRes), nil
}
func (l *nilLog) FindCreateGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	*metrics.FCGraphRes, error,
) {
	return new(metrics.FCGraphRes), nil
}
func (l *nilLog) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	*metrics.FCMetricRes, error,
) {
	return new(metrics.FCMetricRes), nil
}

// NewNopMetricLog returns a no-op metric logger
func newNopMetricLog() metricLogger {
	return &nilLog{}
}
