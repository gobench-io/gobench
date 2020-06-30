package server

import (
	"context"
	"testing"

	"github.com/gobench-io/gobench/ent"
	"github.com/stretchr/testify/assert"
)

func TestNextApplication(t *testing.T) {
	t.Run("empty application", func(t *testing.T) {
		s, _ := NewServer(DefaultMasterOptions())
		// disable the schedule
		s.isSchedule = false
		assert.Nil(t, s.Start())
		assert.Nil(t, s.cleanupDB())

		_, err := s.nextApplication()
		assert.True(t, ent.IsNotFound(err))
	})

	t.Run("empty application", func(t *testing.T) {
		s, _ := NewServer(DefaultMasterOptions())
		// disable the schedule
		s.isSchedule = false
		assert.Nil(t, s.Start())
		assert.Nil(t, s.cleanupDB())

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
}
