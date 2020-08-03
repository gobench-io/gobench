package server

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

// FindCreateGroup find or create new group
// return the existing/new group ent, is created, and error
func (m *master) FindCreateGroup(ctx context.Context, mg metrics.Group, appID int) (
	eg *ent.Group, err error,
) {
	eg, err = m.job.app.
		QueryGroups().
		Where(
			entGroup.NameEQ(mg.Name),
			entGroup.HasApplicationWith(
				entApp.IDEQ(appID),
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
			SetName(mg.Name).
			SetApplicationID(m.job.app.ID).
			Save(ctx)
		return
	}

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
	if err != nil {
		if !ent.IsNotFound(err) {
			return
		}

		egraph, err = m.db.Graph.Create().
			SetTitle(mgraph.Title).
			SetUnit(mgraph.Unit).
			SetGroupID(groupID).
			Save(ctx)
		return
	}
	return
}

func (m *master) FindCreateMetric(ctx context.Context, mmetric metrics.Metric, graphID int) (
	emetric *ent.Metric, err error,
) {
	emetric, err = m.db.Metric.Query().
		Where(
			entMetric.TitleEQ(mmetric.Title),
			entMetric.TypeEQ(string(mmetric.Type)),
			entMetric.HasGraphWith(
				entGraph.IDEQ(graphID),
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
			SetTitle(mmetric.Title).
			SetType(string(mmetric.Type)).
			SetGraphID(graphID).
			Save(ctx)

		return
	}
	return
}

// rpc interface

type FCGroupReq struct {
	Name  string
	AppID int
}
type FCGroupRes struct {
	ID int
}

// FindCreateGroupRPC find or create new group
// return the existing/new group ent, is created, and error
func (m *master) FindCreateGroupRPC(args *FCGroupReq, reply *FCGroupRes) (err error) {
	ctx := context.TODO()

	var eg *ent.Group

	defer func() {
		if err == nil {
			reply.ID = eg.ID
		}
	}()

	eg, err = m.job.app.
		QueryGroups().
		Where(
			entGroup.NameEQ(args.Name),
			entGroup.HasApplicationWith(
				entApp.IDEQ(args.AppID),
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
			SetName(args.Name).
			SetApplicationID(m.job.app.ID).
			Save(ctx)

		return
	}

	return
}

type FCGraphReq struct {
	Title   string
	Unit    string
	GroupID int
}
type FCGraphRes struct {
	ID int
}

func (m *master) FindCreateGraphRPC(req *FCGraphReq, res *FCGraphRes) (err error) {
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
