package master

import (
	"context"
	"testing"
	"time"

	entApplication "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	"github.com/gobench-io/gobench/pb"
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
	assert.Equal(t, g.ID, groupRes.Id)

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
	assert.Equal(t, g.ID, graphRes.Id)

	// call the same RPC, the result should be like before
	graphRes2, err := m.FindCreateGraph(ctx, graphReq)
	assert.Nil(t, err)
	assert.Equal(t, graphRes, graphRes2)
}

// func TestFindCreateMetric(t *testing.T) {
// 	var err error
// 	ctx := context.TODO()

// 	m := seedMaster(t)

// 	m.job.app, err = m.NewApplication(ctx, "name", "scenario", "", "")
// 	assert.Nil(t, err)

// 	prefix := time.Now().String()
// 	groupName := "HTTP (" + prefix + ")"

// 	// create new group
// 	groupRes := new(metrics.FCGroupRes)
// 	assert.Nil(t, m.FindCreateGroup(
// 		&metrics.FCGroupReq{Name: groupName, AppID: m.job.app.ID},
// 		groupRes))

// 	// create new graph
// 	graphReq := &metrics.FCGraphReq{
// 		Title:   "HTTP Response",
// 		Unit:    "N",
// 		GroupID: groupRes.ID,
// 	}
// 	graphRes := new(metrics.FCGraphRes)
// 	assert.Nil(t, m.FindCreateGraph(graphReq, graphRes))

// 	// create new metric
// 	metricReq := &metrics.FCMetricReq{
// 		Title:   ".http_ok",
// 		Type:    metrics.Counter,
// 		GraphID: graphRes.ID,
// 	}
// 	metricRes := new(metrics.FCMetricRes)
// 	assert.Nil(t, m.FindCreateMetric(metricReq, metricRes))

// 	// call the same RPC, the result should be like before
// 	metricRes2 := new(metrics.FCGraphRes)
// 	assert.Nil(t, m.FindCreateGraph(graphReq, metricRes2))
// 	assert.Equal(t, graphRes, metricRes2)

// 	// read from db, check with groupRes
// 	metrics, err := m.db.Metric.Query().Where(
// 		entMetric.TitleEQ(metricReq.Title),
// 		entMetric.HasGraphWith(
// 			entGraph.IDEQ(graphRes.ID),
// 		),
// 	).All(ctx)
// 	assert.Nil(t, err)
// 	assert.Len(t, metrics, 1)
// 	m0 := metrics[0]
// 	assert.Equal(t, m0.ID, metricRes.ID)
// }
