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

func (m *master) Counter(ctx context.Context, mID int, wid, title string, time, c int64) error {
	_, err := m.db.Counter.Create().
		SetMetricID(mID).
		SetCount(c).
		SetTime(time).
		SetWID(wid).
		Save(ctx)
	return err
}

func (m *master) Histogram(ctx context.Context, mID int, wid, title string, time int64, h gometrics.Histogram) error {
	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
	_, err := m.db.Histogram.Create().
		SetMetricID(mID).
		SetCount(h.Count()).
		SetMin(h.Min()).
		SetMax(h.Max()).
		SetMean(h.Mean()).
		SetStddev(h.StdDev()).
		SetMedian(ps[0]).
		SetP75(ps[1]).
		SetP95(ps[2]).
		SetP99(ps[3]).
		SetP999(ps[4]).
		SetTime(time).
		Save(ctx)
	return err
}

func (m *master) Gauge(ctx context.Context, mID int, wid, title string, time int64, g int64) error {
	_, err := m.db.Gauge.Create().
		SetMetricID(mID).
		SetValue(g).
		SetTime(time).
		SetWID(wid).
		Save(ctx)
	return err
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
