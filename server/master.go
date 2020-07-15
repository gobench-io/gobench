package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/worker"
)

// job status. The job is in either pending, provisioning, running, finished
// cancel, error states
type jobState string

// App states
const (
	jobPending      jobState = "pending"
	jobProvisioning jobState = "provisioning"
	jobRunning      jobState = "running"
	jobFinished     jobState = "finished"
	jobCancel       jobState = "cancel"
	jobError        jobState = "error"
)

type master struct {
	addr        string // host name
	port        int    // api port
	clusterPort int    // cluster port

	// database
	dbFilename string
	db         *ent.Client

	ml metriclog

	lw  *worker.Worker // local worker
	job job
}

type job struct {
	app    *ent.Application
	plugin string // plugin path
}

// to is the function to set new state for an application
// save new state to the db
func (m *master) jobTo(ctx context.Context, state jobState) error {
	return m.job.app.Update().
		SetStatus(string(state)).
		Exec(ctx)
}

// setupDb setup the db in the master
func (m *master) setupDb() error {
	filename := m.dbFilename
	client, err := ent.Open(
		"sqlite3",
		filename+"?mode=rwc&cache=shared&&_busy_timeout=9999999&_fk=1")

	if err != nil {
		return fmt.Errorf("failed opening sqlite3 connection: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	m.dbFilename = filename
	m.db = client

	return nil
}

// schedule get a pending application from the db if there is no active job
func (m *master) schedule() {
	for {
		time.Sleep(1 * time.Second)
		m.run()
	}
}

func (m *master) run() {
	ctx, _ := context.WithCancel(context.Background())

	// finding pending application
	app, err := m.nextApplication(ctx)

	if err != nil {
		return
	}

	// create new job from the application
	m.job.app = app

	// change job to provisioning
	m.jobTo(ctx, jobProvisioning)

	if err = m.jobCompile(ctx); err != nil {
		return
	}
	// todo: ditribute the plugin to other worker when run in cloud mode
	// in this phase, the server run in local mode

	// change job to running state
	if err = m.jobTo(ctx, jobRunning); err != nil {
		return
	}
	if err = m.runJob(ctx); err != nil {
		return
	}
}

// provision compiles a scenario to golang plugin, distribute the application to
// worker. Return success when the workers confirm that the plugin is ready
func (m *master) provision() (*ent.Application, error) {
	// compile
	return nil, nil
}

func (m *master) nextApplication(ctx context.Context) (*ent.Application, error) {
	app, err := m.db.
		Application.
		Query().
		Where(
			application.Status(string(jobPending)),
		).
		Order(
			ent.Asc(application.FieldCreatedAt),
		).First(ctx)

	return app, err
}

// jobCompile using go to compile a scenario in plugin build mode
// the result is path to so file
func (m *master) jobCompile(ctx context.Context) error {
	var path string

	scen := m.job.app.Scenario

	// save the scenario to a tmp file
	tmpScenF, err := ioutil.TempFile("", "gobench-scenario-*.go")
	if err != nil {
		return fmt.Errorf("failed creating temp scenario file: %v", err)
	}
	tmpScenName := tmpScenF.Name()

	defer os.Remove(tmpScenName) // cleanup

	_, err = tmpScenF.Write([]byte(scen))
	if err != nil {
		return fmt.Errorf("failed write to scenario file: %v", err)
	}

	if err = tmpScenF.Close(); err != nil {
		return fmt.Errorf("failed close the scenario file: %v", err)
	}

	path = fmt.Sprintf("%s.out", tmpScenName)

	// compile the scenario to a tmp file
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o",
		path, tmpScenName)

	// if out, err := cmd.CombinedOutput(); err != nil {
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed compiling the scenario: %v", err)
	}

	m.job.plugin = path

	return nil
}

// runJob run a application in a job
// by create a local worker
func (m *master) runJob(ctx context.Context) error {
	var err error

	if m.lw, err = worker.NewWorker(&m.ml); err != nil {
		return err
	}

	if err = m.lw.Load(m.job.plugin); err != nil {
		return fmt.Errorf("failed load plugin: %v", err)
	}

	if err = m.lw.Run(); err != nil {
		return err
	}

	return nil
}
