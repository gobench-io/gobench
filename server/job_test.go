package server

import (
	"context"
	"testing"

	"github.com/gobench-io/gobench/ent"
	"github.com/stretchr/testify/assert"
)

func seedServer(t *testing.T) *Server {
	s, _ := NewServer(DefaultMasterOptions())
	// disable the schedule
	s.isSchedule = false
	assert.Nil(t, s.Start())
	assert.Nil(t, s.cleanupDB())

	return s
}

func TestNextApplication(t *testing.T) {
	t.Run("empty application", func(t *testing.T) {
		s := seedServer(t)
		_, err := s.nextApplication()
		assert.True(t, ent.IsNotFound(err))
	})

	t.Run("one application", func(t *testing.T) {
		s := seedServer(t)

		ctx := context.TODO()
		_, err := s.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)

		// the next application is not nil
		a, err := s.nextApplication()
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(appPending))
	})

	t.Run("two applications", func(t *testing.T) {
		s := seedServer(t)

		ctx := context.TODO()
		_, err := s.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)
		_, err = s.NewApplication(ctx, "name 2", "scenario 2")
		assert.Nil(t, err)

		// applications is fifo, the next application is name
		a, err := s.nextApplication()
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(appPending))
	})
}

// func TestCompile(t *testing.T) {
// 	t.Run("valid scenario", func(t *testing.T) {
// 		s, _ := NewServer(DefaultMasterOptions())
// 	})
// }
