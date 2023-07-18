package master

import (
	"context"
	"testing"
	"time"

	entApplication "github.com/gobench-io/gobench/v2/ent/application"
	entGraph "github.com/gobench-io/gobench/v2/ent/graph"
	entGroup "github.com/gobench-io/gobench/v2/ent/group"
	entMetric "github.com/gobench-io/gobench/v2/ent/metric"
	"github.com/gobench-io/gobench/v2/executor/metrics"
	"github.com/gobench-io/gobench/v2/pb"
	"github.com/stretchr/testify/assert"
)

func TestFindCreateGroup(t *testing.T) {
	var err error
	ctx := context.TODO()

	m := seedMaster(t)

	m.job.app, err = m.NewApplication(ctx, "name", "scenario", "", "")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes, err := m.FindCreateGroup(
		ctx,
		&pb.FCGroupReq{
			Name:  groupName,
			AppID: int64(m.job.app.ID),
		},
	)
	assert.Nil(t, err)

	// read from db, check with groupRes
	groups, err := m.db.Group.Query().Where(
		entGroup.Name(groupName),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]
	assert.EqualValues(t, g.ID, groupRes.Id)

	// call the same RPC, the result should be like before
	groupRes2, err := m.FindCreateGroup(
		ctx,
		&pb.FCGroupReq{
			Name:  groupName,
			AppID: int64(m.job.app.ID),
		},
	)
	assert.Equal(t, groupRes, groupRes2)
}

func TestFindCreateGraph(t *testing.T) {
	var err error
	ctx := context.TODO()

	m := seedMaster(t)

	m.job.app, err = m.NewApplication(ctx, "name", "scenario", "", "")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes, err := m.FindCreateGroup(ctx, &pb.FCGroupReq{
		Name:  groupName,
		AppID: int64(m.job.app.ID),
	})
	assert.Nil(t, err)

	graphReq := &pb.FCGraphReq{
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: int64(groupRes.Id),
	}
	graphRes, err := m.FindCreateGraph(ctx, graphReq)
	assert.Nil(t, err)

	// read from db, check with groupRes
	graphs, err := m.db.Graph.Query().Where(
		entGraph.TitleEQ(graphReq.Title),
		entGraph.HasGroupWith(
			entGroup.IDEQ(int(groupRes.Id)),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 1)
	g := graphs[0]
	assert.EqualValues(t, g.ID, graphRes.Id)

	// call the same RPC, the result should be like before
	graphRes2, err := m.FindCreateGraph(ctx, graphReq)
	assert.Nil(t, err)
	assert.Equal(t, graphRes, graphRes2)
}

func TestFindCreateMetric(t *testing.T) {
	var err error
	ctx := context.TODO()

	m := seedMaster(t)

	m.job.app, err = m.NewApplication(ctx, "name", "scenario", "", "")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	// create new group
	groupRes, err := m.FindCreateGroup(ctx, &pb.FCGroupReq{
		AppID: int64(m.job.app.ID),
		Name:  groupName,
	})
	assert.Nil(t, err)

	// create new graph
	graphReq := &pb.FCGraphReq{
		AppID:   int64(m.job.app.ID),
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: int64(groupRes.Id),
	}
	graphRes, err := m.FindCreateGraph(ctx, graphReq)
	assert.Nil(t, err)

	// create new metric
	metricReq := &pb.FCMetricReq{
		AppID:   int64(m.job.app.ID),
		Title:   ".http_ok",
		Type:    string(metrics.Counter),
		GraphID: int64(graphRes.Id),
	}
	metricRes, err := m.FindCreateMetric(ctx, metricReq)
	assert.Nil(t, err)

	// call the same RPC, the result should be like before
	metricRes2, err := m.FindCreateMetric(ctx, metricReq)
	assert.Equal(t, metricRes, metricRes2)

	// read from db, check with groupRes
	metrics, err := m.db.Metric.Query().Where(
		entMetric.TitleEQ(metricReq.Title),
		entMetric.HasGraphWith(
			entGraph.IDEQ(int(graphRes.Id)),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, metrics, 1)
	m0 := metrics[0]
	assert.EqualValues(t, m0.ID, metricRes.Id)
}
