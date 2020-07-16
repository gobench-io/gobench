package server

import (
	"context"

	gometrics "github.com/rcrowley/go-metrics"
)

func (m *master) Counter(ctx context.Context, wid, title string, time, c int64) error {
	m.db.Counter.Create().
		SetCount(c).
		SetTime(time).
		// SetMetricID().
		Save(ctx)
	return nil
}

func (m *master) Histogram(ctx context.Context, wid, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (m *master) Gauge(ctx context.Context, wid, title string, time int64, g int64) error {
	return nil
}
