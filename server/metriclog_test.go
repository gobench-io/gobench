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

	groupReply := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupArgs{Name: groupName, AppID: s.master.job.app.ID},
		groupReply))

	// read from db, check with groupReply
	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name(groupName),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]
	assert.Equal(t, g.ID, groupReply.ID)

	// call the same RPC, the result should be like before
	groupReply2 := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupArgs{Name: groupName, AppID: s.master.job.app.ID},
		groupReply2))
	assert.Equal(t, groupReply, groupReply2)
}

func TestFindCreateGraphRPC(t *testing.T) {
	var err error
	ctx := context.TODO()

	s := seedServer(t)

	s.master.job.app, err = s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	prefix := time.Now().String()
	groupName := "HTTP (" + prefix + ")"

	groupReply := new(FCGroupRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGroupArgs{Name: groupName, AppID: s.master.job.app.ID},
		groupReply))

	// create new graph
	graphReq := &FCGraphReq{
		Title:   "HTTP Response",
		Unit:    "N",
		GroupID: groupReply.ID,
	}
	graphReply := new(FCGraphRes)
	assert.Nil(t, s.master.FindCreateGraphRPC(graphReq, graphReply))

	// read from db, check with groupReply
	graphs, err := s.master.db.Graph.Query().Where(
		entGraph.TitleEQ(graphReq.Title),
		entGraph.HasGroupWith(
			entGroup.IDEQ(groupReply.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 1)
	g := graphs[0]
	assert.Equal(t, g.ID, graphReply.ID)
}
