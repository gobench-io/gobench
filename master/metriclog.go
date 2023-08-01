package master

import (
	"context"

	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/pb"

	entApp "github.com/gobench-io/gobench/v2/ent/application"
	entGraph "github.com/gobench-io/gobench/v2/ent/graph"
	entGroup "github.com/gobench-io/gobench/v2/ent/group"
	entMetric "github.com/gobench-io/gobench/v2/ent/metric"
)

func (m *Master) Counter(ctx context.Context, req *pb.CounterReq) (*pb.CounterRes, error) {
	// todo: check appID condition
	_, err := m.db.Counter.Create().
		SetWID(req.Base.EID).
		SetMetricID(int(req.Base.MID)).
		SetTime(req.Base.Time).
		SetCount(req.Count).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	res := new(pb.CounterRes)

	return res, nil
}

func (m *Master) Histogram(ctx context.Context, req *pb.HistogramReq) (*pb.HistogramRes, error) {
	// todo: check appID condition
	_, err := m.db.Histogram.Create().
		SetWID(req.Base.EID).
		SetMetricID(int(req.Base.MID)).
		SetTime(req.Base.Time).
		SetCount(req.Histogram.Count).
		SetMin(req.Histogram.Min).
		SetMax(req.Histogram.Max).
		SetMean(req.Histogram.Mean).
		SetStddev(req.Histogram.Stddev).
		SetMedian(req.Histogram.Median).
		SetP75(req.Histogram.P75).
		SetP95(req.Histogram.P95).
		SetP99(req.Histogram.P99).
		SetP999(req.Histogram.P999).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	res := new(pb.HistogramRes)

	return res, nil
}

func (m *Master) Gauge(ctx context.Context, req *pb.GaugeReq) (*pb.GaugeRes, error) {
	// todo: check appID condition
	_, err := m.db.Gauge.Create().
		SetWID(req.Base.EID).
		SetMetricID(int(req.Base.MID)).
		SetTime(req.Base.Time).
		SetValue(req.Gauge).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	res := new(pb.GaugeRes)

	return res, nil
}

// FindCreateGroup find or create new group
// return the existing/new group ent, is created, and error
func (m *Master) FindCreateGroup(ctx context.Context, req *pb.FCGroupReq) (res *pb.FCGroupRes, err error) {
	var eg *ent.Group
	res = new(pb.FCGroupRes)

	defer func() {
		if err == nil {
			res.Id = int64(eg.ID)
		}
	}()

	eg, err = m.job.app.
		QueryGroups().
		Where(
			entGroup.NameEQ(req.Name),
			entGroup.HasApplicationWith(
				entApp.IDEQ(int(req.AppID)),
			),
		).
		First(ctx)

	// if there is one found
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
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

func (m *Master) FindCreateGraph(ctx context.Context, req *pb.FCGraphReq) (res *pb.FCGraphRes, err error) {
	var egraph *ent.Graph
	res = new(pb.FCGraphRes)

	defer func() {
		if err == nil {
			res.Id = int64(egraph.ID)
		}
	}()

	egraph, err = m.db.Graph.Query().
		Where(
			entGraph.TitleEQ(req.Title),
			entGraph.UnitEQ(req.Unit),
			entGraph.HasGroupWith(
				entGroup.IDEQ(int(req.GroupID)),
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
			SetGroupID(int(req.GroupID)).
			Save(ctx)
		return
	}

	return
}

func (m *Master) FindCreateMetric(ctx context.Context, req *pb.FCMetricReq) (res *pb.FCMetricRes, err error) {
	var emetric *ent.Metric
	res = new(pb.FCMetricRes)

	defer func() {
		if err == nil {
			res.Id = int64(emetric.ID)
		}
	}()

	emetric, err = m.db.Metric.Query().
		Where(
			entMetric.TitleEQ(req.Title),
			entMetric.TypeEQ(string(req.Type)),
			entMetric.HasGraphWith(
				entGraph.IDEQ(int(req.GraphID)),
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
			SetGraphID(int(req.GraphID)).
			Save(ctx)

		return
	}
	return
}
