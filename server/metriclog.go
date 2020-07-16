package server

import (
	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"

	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"
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

func (m *master) NewGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	egraph *ent.Graph, err error,
) {
	egraph, err = m.db.Graph.Create().
		SetTitle(mgraph.Title).
		SetUnit(mgraph.Unit).
		SetGroupID(groupID).
		Save(ctx)
	return
}

func (m *master) NewMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	emetric *ent.Metric, created bool, err error,
) {
	emetric, err = m.db.Metric.Query().
		Where(
			entMetric.TitleEQ(mmetric.Title),
		).
		First(ctx)

	// if there is one found
	if err != nil && !ent.IsNotFound(err) {
		return
	}

	emetric, err = m.db.Metric.
		Create().
		SetTitle(mmetric.Title).
		SetType(string(mmetric.Type)).
		SetGraphID(graphID).
		Save(ctx)

	if err != nil {
		return
	}

	created = true

	return
}
