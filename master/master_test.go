package server

import (
	"context"
	"testing"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/stretchr/testify/assert"
)

func seedServer(t *testing.T) *Server {
	log := logger.NewStdLogger()

	var err error
	s, _ := NewServer(DefaultMasterOptions())
	// disable the schedule
	s.isSchedule = false
	s.master.logger = log
	s.master.job = &job{}
	s.master.job.app = &ent.Application{}
	assert.Nil(t, err)
	assert.Nil(t, s.Start())

	return s
}

func TestNextApplication(t *testing.T) {
	t.Run("empty application", func(t *testing.T) {
		ctx := context.Background()
		s := seedServer(t)
		assert.Nil(t, s.cleanupDB())
		_, err := s.master.nextApplication(ctx)
		assert.True(t, ent.IsNotFound(err))
	})

	t.Run("one application", func(t *testing.T) {
		ctx := context.Background()
		s := seedServer(t)
		assert.Nil(t, s.cleanupDB())

		_, err := s.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)

		// the next application is not nil
		a, err := s.master.nextApplication(ctx)
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(jobPending))
	})

	t.Run("two applications", func(t *testing.T) {
		ctx := context.Background()
		s := seedServer(t)

		_, err := s.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)
		_, err = s.NewApplication(ctx, "name 2", "scenario 2")
		assert.Nil(t, err)

		// applications is fifo, the next application is name
		a, err := s.master.nextApplication(ctx)
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(jobPending))
	})
}

func TestCompile(t *testing.T) {
	t.Run("invalid scenario", func(t *testing.T) {
		ctx := context.Background()

		s := seedServer(t)
		s.master.job.app.Scenario = `
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
// missing f1 function`

		err := s.master.jobCompile(ctx)
		assert.EqualError(t, err, "failed compiling the scenario: exit status 2")
		assert.NoFileExists(t, s.master.job.plugin)
	})

	t.Run("valid scenario", func(t *testing.T) {
		ctx := context.Background()
		s := seedServer(t)
		s.master.job.app.Scenario = `
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
		err := s.master.jobCompile(ctx)
		assert.Nil(t, err)
		assert.FileExists(t, s.master.job.plugin)
	})
}

func TestRun(t *testing.T) {
	ctx := context.Background()
	s := seedServer(t)
	s.master.job.app.Scenario = `
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
			Nu:   1,
			Rate: 100,
			Fu:   f1,
		},
	}
}

func f1(ctx context.Context, vui int) {
	time.Sleep(1 * time.Second)
	log.Println("tic")
}`

	err := s.master.jobCompile(ctx)
	assert.Nil(t, err)
	// should run for mor than 1 seconds
	assert.Nil(t, s.master.runJob(ctx))
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	s := seedServer(t)

	scenario := `
package main

import (
	"context"

	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   1,
			Rate: 100,
			Fu:   f1,
		},
	}
}

func f1(ctx context.Context, vui int) {
	for {}
}`

	app, _ := s.NewApplication(ctx, "cancel test", scenario)
	j := &job{
		app:    app,
		cancel: cancel,
	}

	go func() {
		time.Sleep(1 * time.Second)
		assert.Equal(t, string(jobRunning), s.master.job.app.Status, "should run after 1 second")

		assert.Nil(t, s.master.cancel(ctx, app.ID))
	}()

	err := s.master.run(ctx, j)
	assert.EqualError(t, err, ErrAppIsCanceled.Error())
}

func TestMetricLogSetup(t *testing.T) {
	ctx := context.Background()
	s := seedServer(t)

	scenario := `
package main

import (
	"context"

	httpClient "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   1,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	_, err := httpClient.NewHttpClient(ctx, "home")
	if err != nil {
		panic(err)
	}
}
`

	app, _ := s.NewApplication(ctx, "http metric log setup test", scenario)
	j := &job{
		app: app,
	}

	err := s.master.run(ctx, j)
	assert.Nil(t, err)
}
