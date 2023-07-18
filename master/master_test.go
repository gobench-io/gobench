package master

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/logger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func seedMaster(t *testing.T) *Master {
	logger := logger.NewStdLogger()
	m, err := NewMaster(&Options{
		Addr:    "0.0.0.0",
		Port:    8080,
		HomeDir: "/tmp",
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

func (m *Master) seedApplication(ctx context.Context, t *testing.T) *ent.Application {
	app, err := m.NewApplication(ctx, "foo", "bar", "", "")
	assert.Nil(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, app.Name, "foo")
	assert.Equal(t, app.Scenario, "bar")
	return app
}

func localGobenchMod(t *testing.T) string {
	testDir, _ := os.Getwd()
	mainDir, _ := exec.Command("dirname", testDir).CombinedOutput()
	return fmt.Sprintf(`
		module gobench.io/scenario
		replace github.com/gobench-io/gobench/v2 => %s
		`, string(mainDir))
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

		_, err := m.NewApplication(ctx, "name", "scenario", "", "")
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

		_, err := m.NewApplication(ctx, "name", "scenario", "", "")
		assert.Nil(t, err)
		_, err = m.NewApplication(ctx, "name 2", "scenario 2", "", "")
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
// export is a required function for a scenario
func export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   20,
			Rate: 100,
			Fu:   f1,
		},
	}
}
`
		m.job.logger = logger.NewNopLogger()

		err := m.jobCompile(ctx)
		assert.EqualError(t, err, "compile scenario: exit status 1")
		assert.NoFileExists(t, m.job.plugin)
	})

	t.Run("valid scenario", func(t *testing.T) {
		ctx := context.Background()
		m := seedMaster(t)

		m.job.app.Gomod = localGobenchMod(t)
		m.job.app.Scenario = `
package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/executor/scenario"
)

// export is a required function for a scenario
func export() scenario.Vus {
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
		m.job.logger = logger.NewNopLogger()

		err := m.jobCompile(ctx)
		assert.Nil(t, err)
		assert.FileExists(t, m.job.plugin)
	})
}

func TestRun(t *testing.T) {
	ctx := context.Background()
	m := seedMaster(t)
	gomod := localGobenchMod(t)
	scenario := `
package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/executor/scenario"
)

// export is a required function for a scenario
func export() scenario.Vus {
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
	app, err := m.NewApplication(ctx, "test run", scenario, gomod, "")
	assert.Nil(t, err)

	m.job = &job{
		app: app,
	}
	_, err = m.job.setLogs(m.Logpaths(app.ID))
	assert.Nil(t, err)

	assert.Nil(t, m.jobCompile(ctx))

	// should run for mor than 1 seconds
	assert.Nil(t, m.runJob(ctx))
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	m := seedMaster(t)

	gomod := localGobenchMod(t)
	scenario := `
package main

import (
	"context"

	"github.com/gobench-io/gobench/v2/executor/scenario"
)

func export() scenario.Vus {
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

	app, _ := m.NewApplication(ctx, "cancel test", scenario, gomod, "")
	m.job = &job{
		app:    app,
		cancel: cancel,
	}

	_, err := m.job.setLogs(m.Logpaths(app.ID))
	assert.Nil(t, err)

	go func() {
		count := 0
		for {
			time.Sleep(1 * time.Second)
			count++
			if string(jobRunning) != m.job.app.Status && count <= 5 {
				continue
			}
			break
		}
		assert.Equal(t, string(jobRunning), m.job.app.Status, "should run after 5 second")

		assert.Nil(t, m.cancel(ctx, app.ID))
	}()

	err = m.run(ctx)
	assert.EqualError(t, err, ErrAppIsCanceled.Error())
}

func TestMetricLogSetup(t *testing.T) {
	ctx := context.Background()
	m := seedMaster(t)

	gomod := localGobenchMod(t)
	scenario := `
package main

import (
	"context"

	httpClient "github.com/gobench-io/gobench/v2/clients/http"
	"github.com/gobench-io/gobench/v2/executor/scenario"
)

func export() scenario.Vus {
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

	app, _ := m.NewApplication(ctx, "http metric log setup test", scenario, gomod, "")
	m.job = &job{
		app: app,
	}

	_, err := m.job.setLogs(m.Logpaths(app.ID))
	assert.Nil(t, err)

	err = m.run(ctx)
	assert.Nil(t, err)
}

func TestLogpaths(t *testing.T) {
	m := seedMaster(t)

	folder, sf, uf := m.Logpaths(12)

	assert.Contains(t, folder, "/tmp/applications/12")
	assert.Contains(t, sf, "/tmp/applications/12/system.log")
	assert.Contains(t, uf, "/tmp/applications/12/user.log")
}

func TestMaster_SetApplicationTag(t *testing.T) {
	m := seedMaster(t)
	ctx := context.Background()
	app := m.seedApplication(ctx, t)

	type args struct {
		appID int
		tag   string
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Tag
		wantErr bool
	}{
		{
			name:    "should return error with missing appID and tag",
			args:    args{},
			want:    nil,
			wantErr: true,
		},

		{
			name: "should return error with empty tag name",
			args: args{
				tag:   "",
				appID: app.ID,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return error with missing appID",
			args: args{
				tag: "bar",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return success with valid name",
			args: args{
				tag:   "foo",
				appID: app.ID,
			},
			want: &ent.Tag{
				Name: "foo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := m.SetApplicationTag(ctx, tt.args.appID, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Master.SetApplicationTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Master.SetApplicationTag() = %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}

func TestMaster_GetTagByApplication(t *testing.T) {
	m := seedMaster(t)
	ctx := context.Background()
	tagName := "foo"
	app := m.seedApplication(ctx, t)
	tag, err := m.SetApplicationTag(ctx, app.ID, tagName)
	if err != nil {
		t.Fatal("Master.GetTagByApplication() seed SetApplicationTag error")
	}

	type args struct {
		app     *ent.Application
		tagName string
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Tag
		wantErr bool
	}{
		{
			name: "should return success with valid request",
			args: args{
				tagName: tag.Name,
				app:     app,
			},
			want:    tag,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetTagByApplication(ctx, tt.args.app, tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Master.GetTagByApplication) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Master.GetTagByApplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_GetTagByID(t *testing.T) {
	m := seedMaster(t)
	ctx := context.Background()
	tagName := "foo"
	app := m.seedApplication(ctx, t)
	tag, err := m.SetApplicationTag(ctx, app.ID, tagName)
	if err != nil {
		t.Fatal("Master.GetTagByApplication() seed SetApplicationTag error")
	}
	type args struct {
		tagID int
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Tag
		wantErr bool
	}{
		{
			name: "should return success with valid request",
			args: args{
				tagID: tag.ID,
			},
			want:    tag,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.GetTagByID(ctx, tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Master.GetTagByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Master.GetTagByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_RemoveApplicationTag(t *testing.T) {
	m := seedMaster(t)
	ctx := context.Background()
	tagName := "foo"
	app := m.seedApplication(ctx, t)
	tag, err := m.SetApplicationTag(ctx, app.ID, tagName)
	if err != nil {
		t.Fatal("Master.GetTagByApplication() seed SetApplicationTag error")
	}

	type args struct {
		tag *ent.Tag
	}
	tests := []struct {
		name    string
		args    args
		want    *ent.Tag
		wantErr bool
	}{
		{
			name: "should return success with valid request",
			args: args{
				tag: tag,
			},
			want:    tag,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.RemoveApplicationTag(ctx, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("Master.RemoveApplicationTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
