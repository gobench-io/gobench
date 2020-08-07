package master

import (
	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"
	gometrics "github.com/rcrowley/go-metrics"

	entApp "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"
)

func (m *master) Counter(ctx context.Context, mID int, wid, title string, time, c int64) error {
	_, err := m.db.Counter.Create().
		SetWID(wid).
		SetMetricID(mID).
		SetTime(time).
		SetCount(c).
		Save(ctx)
	return err
}

func (m *master) Histogram(ctx context.Context, mID int, wid, title string, time int64, h gometrics.Histogram) error {
	ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
	_, err := m.db.Histogram.Create().
		SetWID(wid).
		SetMetricID(mID).
		SetTime(time).
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
		Save(ctx)
	return err
}

func (m *master) Gauge(ctx context.Context, mID int, wid, title string, time int64, g int64) error {
	_, err := m.db.Gauge.Create().
		SetWID(wid).
		SetMetricID(mID).
		SetTime(time).
		SetValue(g).
		Save(ctx)
	return err
}

// rpc interface

// FindCreateGroupRPC find or create new group
// return the existing/new group ent, is created, and error
func (m *master) FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) (err error) {
	ctx := context.TODO()

	var eg *ent.Group

	defer func() {
		if err == nil {
			res.ID = eg.ID
		}
	}()

	eg, err = m.job.app.
		QueryGroups().
		Where(
			entGroup.NameEQ(req.Name),
			entGroup.HasApplicationWith(
				entApp.IDEQ(req.AppID),
			),
		).
		First(ctx)

	// if there is one found
	if err != nil {
		if !ent.IsNotFound(err) {
			return
		}

		eg, err = m.db.Group.
			Create().
			SetName(req.Name).
			SetApplicationID(m.job.app.ID).
			Save(ctx)

		return
	}

	return
}

func (m *master) FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) (err error) {
	ctx := context.TODO()

	var egraph *ent.Graph

	defer func() {
		if err == nil {
			res.ID = egraph.ID
		}
	}()

	egraph, err = m.db.Graph.Query().
		Where(
			entGraph.TitleEQ(req.Title),
			entGraph.UnitEQ(req.Unit),
			entGraph.HasGroupWith(
				entGroup.IDEQ(req.GroupID),
			),
		).
		First(ctx)

	// if there is one found
	if err != nil {
		if !ent.IsNotFound(err) {
			return
		}

		egraph, err = m.db.Graph.Create().
			SetTitle(req.Title).
			SetUnit(req.Unit).
			SetGroupID(req.GroupID).
			Save(ctx)
		return
	}

	return
}

func (m *master) FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) (err error) {
	ctx := context.TODO()

	var emetric *ent.Metric

	defer func() {
		if err == nil {
			res.ID = emetric.ID
		}
	}()

	emetric, err = m.db.Metric.Query().
		Where(
			entMetric.TitleEQ(req.Title),
			entMetric.TypeEQ(string(req.Type)),
			entMetric.HasGraphWith(
				entGraph.IDEQ(req.GraphID),
			),
		).
		First(ctx)

	// if there is one found
	if err != nil {
		if !ent.IsNotFound(err) {
			return
		}

		emetric, err = m.db.Metric.
			Create().
			SetTitle(req.Title).
			SetType(string(req.Type)).
			SetGraphID(req.GraphID).
			Save(ctx)

		return
	}
	return
}
