package server

import (
	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"

	entGraph "github.com/gobench-io/gobench/ent/graph"
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

// FindCreateGroup find or create new group
// return the existing/new group ent, is created, and error
func (m *master) FindCreateGroup(ctx context.Context, mg metrics.Group) (
	eg *ent.Group, err error,
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

	return
}

func (m *master) FindCreateGraph(ctx context.Context, mgraph metrics.Graph, groupID int) (
	egraph *ent.Graph, err error,
) {
	egraph, err = m.db.Graph.Query().
		Where(
			entGraph.TitleEQ(mgraph.Title),
			entGraph.UnitEQ(mgraph.Unit),
			entGraph.HasGroupWith(
				entGroup.IDEQ(groupID),
			),
		).
		First(ctx)

	// if there is one found
	if err != nil && !ent.IsNotFound(err) {
		return
	}

	egraph, err = m.db.Graph.Create().
		SetTitle(mgraph.Title).
		SetUnit(mgraph.Unit).
		SetGroupID(groupID).
		Save(ctx)
	return
}

func (m *master) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	emetric *ent.Metric, err error,
) {
	emetric, err = m.db.Metric.Query().
		Where(
			entMetric.TitleEQ(mmetric.Title),
			entMetric.TypeEQ(mmetric.Title),
			entMetric.HasGraphWith(
				entGraph.IDEQ(graphID),
			),
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

	return
}
