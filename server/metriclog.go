package server

import (
	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"

	entGroup "github.com/gobench-io/gobench/ent/group"
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

// NewGroup find or create new group
// return the existing/new group ent, is created, and error
func (m *master) NewGroup(ctx context.Context, mg metrics.Group) (
	eg *ent.Group, created bool, err error,
) {
	eg, err = m.job.app.
		QueryGroups().
		Where(
			entGroup.NameEQ(mg.Name),
		).
		First(ctx)

	// if there is one found
	if err != nil && !ent.IsNotFound(err) {
		return
	}

	eg, err = m.db.Group.
		Create().
		SetName(mg.Name).
		SetApplicationID(m.job.app.ID).
		Save(ctx)

	if err != nil {
		return
	}

	created = true

	return eg, created, err
}
