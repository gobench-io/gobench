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

func TestCompile(t *testing.T) {
	t.Run("valid scenario", func(t *testing.T) {
		s := seedServer(t)
		scen := `
package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/scenario"
)

// Export is a required function for a scenario
func Export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   20,
			Rate: 100,
			Fu:   f1,
		},
	}
}

// this function receive the ctx.Done signal
func f1(ctx context.Context, vui int) {
	for {
		log.Println("tic")
		time.Sleep(1 * time.Second)
	}
}`
		path, err := s.compile(scen)
		assert.Nil(t, err)
		assert.FileExists(t, path)
	})
}
