package server

import (
	"context"
	"testing"
	"time"

	entApplication "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	"github.com/stretchr/testify/assert"
)

func TestFindCreateGroupRPC(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes))

	// read from db, check with groupRes
	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name(groupName),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]
	assert.Equal(t, g.ID, groupRes.ID)

	// call the same RPC, the result should be like before
	groupRes2 := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes2))
	assert.Equal(t, groupRes, groupRes2)
}

func TestFindCreateGraphRPC(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupRes := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupReq{Name: groupName, AppID: s.master.job.app.ID},
		groupRes))

	// create new graph
	graphReq := &FCGraphReq{
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: groupRes.ID,
	}
	graphRes := new(FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraphRPC(graphReq, graphRes))

	// read from db, check with groupRes
	graphs, err := s.master.db.Graph.Query().Where(
		entGraph.TitleEQ(graphReq.Title),
		entGraph.HasGroupWith(
			entGroup.IDEQ(groupRes.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 1)
	g := graphs[0]
	assert.Equal(t, g.ID, graphRes.ID)

	// call the same RPC, the result should be like before
	graphRes2 := new(FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraphRPC(graphReq, graphRes2))
	assert.Equal(t, graphRes, graphRes2)
}
