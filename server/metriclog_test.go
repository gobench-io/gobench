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
	ctx := context.TODO()

	s := seedServer(t)

	_, err := s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	s.master.job.app, err = s.master.nextApplication(ctx)
	assert.Nil(t, err)

	prefix := time.Now().String()
	name := "HTTP (" + prefix + ")"

	reply := new(FCGRes)

	err = s.master.FindCreateGroupRPC(&FCGArgs{Name: name}, reply)
	assert.Nil(t, err)

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
}
