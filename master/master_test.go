package master

import (
	"context"
	"testing"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/stretchr/testify/assert"
)

func seedMaster(t *testing.T) *Master {
	logger := logger.NewStdLogger()
	m, err := NewMaster(&Options{
		Addr:    "0.0.0.0",
		Port:    8080,
		DbPath:  "./gobench.sqlite3",
		Program: "gobench",
	}, logger)

	// disable the schedule
	m.isScheduled = false
	m.job = &job{}
	m.job.app = &ent.Application{}
	assert.Nil(t, err)
	assert.Nil(t, m.Start())

	return m
}

func TestNextApplication(t *testing.T) {
	t.Run("empty application", func(t *testing.T) {
		ctx := context.Background()
		m := seedMaster(t)
		assert.Nil(t, m.cleanupDB())
		_, err := m.nextApplication(ctx)
		assert.True(t, ent.IsNotFound(err))
	})

	t.Run("one application", func(t *testing.T) {
		ctx := context.Background()
		m := seedMaster(t)
		assert.Nil(t, m.cleanupDB())

		_, err := m.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)

		// the next application is not nil
		a, err := m.nextApplication(ctx)
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(jobPending))
	})

	t.Run("two applications", func(t *testing.T) {
		ctx := context.Background()
		m := seedMaster(t)

		_, err := m.NewApplication(ctx, "name", "scenario")
		assert.Nil(t, err)
		_, err = m.NewApplication(ctx, "name 2", "scenario 2")
		assert.Nil(t, err)

		// applications is fifo, the next application is name
		a, err := m.nextApplication(ctx)
		assert.Nil(t, err)
		assert.Equal(t, a.Name, "name")
		assert.Equal(t, a.Scenario, "scenario")
		assert.Equal(t, a.Status, string(jobPending))
	})
}

func TestCompile(t *testing.T) {
	t.Run("invalid scenario", func(t *testing.T) {
		ctx := context.Background()

		m := seedMaster(t)
		m.job.app.Scenario = `
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

		err := m.jobCompile(ctx)
		assert.EqualError(t, err, "failed compiling the scenario: exit status 2")
		assert.NoFileExists(t, m.job.plugin)
	})

	t.Run("valid scenario", func(t *testing.T) {
		ctx := context.Background()
		m := seedMaster(t)
		m.job.app.Scenario = `
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
		err := m.jobCompile(ctx)
		assert.Nil(t, err)
		assert.FileExists(t, m.job.plugin)
	})
}

func TestRun(t *testing.T) {
	ctx := context.Background()
	m := seedMaster(t)
	scenario := `
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
	app, err := m.NewApplication(ctx, "test run", scenario)
	assert.Nil(t, err)

	m.job = &job{
		app: app,
	}

	assert.Nil(t, m.jobCompile(ctx))

	// should run for mor than 1 seconds
	assert.Nil(t, m.runJob(ctx))
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	m := seedMaster(t)

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

	app, _ := m.NewApplication(ctx, "cancel test", scenario)
	j := &job{
		app:    app,
		cancel: cancel,
	}

	go func() {
		time.Sleep(1 * time.Second)
		assert.Equal(t, string(jobRunning), m.job.app.Status, "should run after 1 second")

		assert.Nil(t, m.cancel(ctx, app.ID))
	}()

	err := m.run(ctx, j)
	assert.EqualError(t, err, ErrAppIsCanceled.Error())
}

func TestMetricLogSetup(t *testing.T) {
	ctx := context.Background()
	m := seedMaster(t)

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

	app, _ := m.NewApplication(ctx, "http metric log setup test", scenario)
	j := &job{
		app: app,
	}

	err := m.run(ctx, j)
	assert.Nil(t, err)
}
