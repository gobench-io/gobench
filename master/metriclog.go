package master

import (
	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/metrics"

	entApp "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"
)

func (m *Master) Counter(req *metrics.CounterReq, res *metrics.CounterRes) (
	err error,
) {
	// todo: check appID condition
	ctx := context.TODO()

	_, err = m.db.Counter.Create().
		SetWID(req.EID).
		SetMetricID(req.MID).
		SetTime(req.Time).
		SetCount(req.Count).
		Save(ctx)

	res.AppID = m.job.app.ID
	res.Success = true

	return
}

func (m *Master) Histogram(req *metrics.HistogramReq, res *metrics.HistogramRes) (
	err error,
) {
	// todo: check appID condition
	ctx := context.TODO()

	_, err = m.db.Histogram.Create().
		SetWID(req.EID).
		SetMetricID(req.MID).
		SetTime(req.Time).
		SetCount(req.Count).
		SetMin(req.Min).
		SetMax(req.Max).
		SetMean(req.Mean).
		SetStddev(req.Stddev).
		SetMedian(req.Median).
		SetP75(req.P75).
		SetP95(req.P95).
		SetP99(req.P99).
		SetP999(req.P999).
		Save(ctx)

	res.AppID = m.job.app.ID
	res.Success = true

	return
}

func (m *Master) Gauge(req *metrics.GaugeReq, res *metrics.GaugeRes) (
	err error,
) {
	// todo: check appID condition
	ctx := context.TODO()

	_, err = m.db.Gauge.Create().
		SetWID(req.EID).
		SetMetricID(req.MID).
		SetTime(req.Time).
		SetValue(req.Gauge).
		Save(ctx)

	res.AppID = m.job.app.ID
	res.Success = true

	return
}

// FindCreateGroup find or create new group
// return the existing/new group ent, is created, and error
func (m *Master) FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) (err error) {
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

func (m *Master) FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) (err error) {
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

func (m *Master) FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) (err error) {
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
