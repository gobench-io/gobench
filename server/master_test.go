package server

import (
	"context"
	"testing"
	"time"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/worker"
	"github.com/stretchr/testify/assert"

	entApplication "github.com/gobench-io/gobench/ent/application"
	entGraph "github.com/gobench-io/gobench/ent/graph"
	entGroup "github.com/gobench-io/gobench/ent/group"
	entMetric "github.com/gobench-io/gobench/ent/metric"
)

func seedServer(t *testing.T) *Server {
	log := logger.NewNopLogger()

	var err error
	s, _ := NewServer(DefaultMasterOptions())
	// disable the schedule
	s.isSchedule = false
	s.master.job = &job{}
	s.master.job.app = &ent.Application{}
	s.master.lw, err = worker.NewWorker(&s.master, log, s.master.job.app.ID)
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

// this function receive the ctx.Done signal
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
	assert.EqualError(t, err, worker.ErrAppCancel.Error())
}

// worker.Setup should create associated tables in the database
func TestSetup(t *testing.T) {
	ctx := context.TODO()

	s := seedServer(t)
	// assert.Nil(t, s.cleanupDB())

	_, err := s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	s.master.job.app, err = s.master.nextApplication(ctx)
	assert.Nil(t, err)

	prefix := time.Now().String()
	group := metrics.Group{
		Name: "HTTP (" + prefix + ")",
		Graphs: []metrics.Graph{
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".http_ok",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_fail",
						Type:  metrics.Counter,
					},
					{
						Title: prefix + ".http_other_fail",
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "Latency",
				Unit:  "Microsecond",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".latency",
						Type:  metrics.Histogram,
					},
				},
			},
		},
	}
	err = worker.Setup([]metrics.Group{
		group,
	})
	assert.Nil(t, err)

	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name("HTTP ("+prefix+")"),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]

	graphs, err := s.master.db.Graph.Query().Where(
		entGraph.HasGroupWith(
			entGroup.IDEQ(g.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 2)

	assert.Equal(t, "HTTP Response", graphs[0].Title)
	assert.Equal(t, "Latency", graphs[1].Title)

	metrics1, err := s.master.db.Metric.Query().Where(
		entMetric.HasGraphWith(
			entGraph.IDEQ(graphs[0].ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, metrics1, 3)
	assert.Equal(t, prefix+".http_ok", metrics1[0].Title)
	assert.Equal(t, prefix+".http_fail", metrics1[1].Title)
	assert.Equal(t, prefix+".http_other_fail", metrics1[2].Title)

	metrics2, err := s.master.db.Metric.Query().Where(
		entMetric.HasGraphWith(
			entGraph.IDEQ(graphs[1].ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, metrics2, 1)
	assert.Equal(t, prefix+".latency", metrics2[0].Title)
}

func TestDucplicateSetup(t *testing.T) {
	ctx := context.TODO()

	s := seedServer(t)
	// assert.Nil(t, s.cleanupDB())

	_, err := s.NewApplication(ctx, "name", "scenario")
	assert.Nil(t, err)

	s.master.job.app, err = s.master.nextApplication(ctx)
	assert.Nil(t, err)

	prefix := time.Now().String()
	group := metrics.Group{
		Name: "HTTP (" + prefix + ")",
		Graphs: []metrics.Graph{
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".http_ok",
						Type:  metrics.Counter,
					},
				},
			},
			{
				Title: "HTTP Response",
				Unit:  "N",
				Metrics: []metrics.Metric{
					{
						Title: prefix + ".http_ok",
						Type:  metrics.Counter,
					},
				},
			},
		},
	}
	err = worker.Setup([]metrics.Group{
		group,
	})
	assert.Nil(t, err)

	groups, err := s.master.db.Group.Query().Where(
		entGroup.Name("HTTP ("+prefix+")"),
		entGroup.HasApplicationWith(
			entApplication.NameEQ("name"),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, groups, 1)
	g := groups[0]

	graphs, err := s.master.db.Graph.Query().Where(
		entGraph.HasGroupWith(
			entGroup.IDEQ(g.ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, graphs, 1)

	metrics, err := s.master.db.Metric.Query().Where(
		entMetric.HasGraphWith(
			entGraph.IDEQ(graphs[0].ID),
		),
	).All(ctx)
	assert.Nil(t, err)
	assert.Len(t, metrics, 1)
}
