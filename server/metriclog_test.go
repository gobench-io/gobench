package server

import (
	"context"
	"testing"
	"time"

	entApplication "github.com/gobench-io/gobench/ent/application"
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
	name := "HTTP (" + prefix + ")"

	reply := new(FCGRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGArgs{Name: name, AppID: s.master.job.app.ID},
		reply))

	// read from db, check with reply
	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name(name),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]
	assert.Equal(t, g.ID, reply.ID)

	// call the same RPC, the result should be like before
	reply2 := new(FCGRes)
	assert.Nil(t, s.master.FindCreateGroupRPC(
		&FCGArgs{Name: name, AppID: s.master.job.app.ID},
		reply2))
	assert.Equal(t, reply, reply2)
}
