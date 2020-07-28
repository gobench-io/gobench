package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"time"

	"context"

	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/worker"
	"go.uber.org/zap"
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
	mu          sync.Mutex
	addr        string // host name
	port        int    // api port
	clusterPort int    // cluster port

	logger *zap.SugaredLogger

	// database
	dbFilename string
	db         *ent.Client

	lw  *worker.Worker // local worker
	job *job
}

type job struct {
	app    *ent.Application
	plugin string // plugin path
	cancel context.CancelFunc
}

// to is the function to set new state for an application
// save new state to the db
func (m *master) jobTo(ctx context.Context, state jobState) (err error) {
	m.job.app, err = m.job.app.Update().
		SetStatus(string(state)).
		Save(ctx)

	return
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
		ctx, cancel := context.WithCancel(context.Background())
		time.Sleep(1 * time.Second)

		// finding pending application
		app, err := m.nextApplication(ctx)
		if err != nil {
			continue
		}
		job := &job{
			app:    app,
			cancel: cancel,
		}
		m.run(ctx, job)
	}
}

func (m *master) run(ctx context.Context, j *job) (err error) {
	// create new job from the application
	m.job = j

	defer func() {
		if err != nil {
			m.logger.Errorw("failed run job",
				"application id", m.job.app.ID,
				"err", err,
			)
			je := jobError
			if errors.Is(err, worker.ErrAppCancel) {
				je = jobCancel
			}
			ctx := context.TODO()
			m.jobTo(ctx, je)
		}
	}()

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	// change job to provisioning
	if err = m.jobTo(ctx, jobProvisioning); err != nil {
		return
	}

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	if err = m.jobCompile(ctx); err != nil {
		return
	}
	// todo: ditribute the plugin to other worker when run in cloud mode
	// in this phase, the server run in local mode

	// change job to running state
	if err = m.jobTo(ctx, jobRunning); err != nil {
		return
	}

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	if err = m.runJob(ctx); err != nil {
		return
	}

	if err = m.jobTo(ctx, jobFinished); err != nil {
		return
	}

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	return
}

// cancel terminates a running job with the same app ID
func (m *master) cancel(ctx context.Context, appID int) error {
	if m.job == nil {
		return ErrAppNotRunning
	}
	if m.job.app.ID != appID {
		return ErrAppNotRunning
	}

	m.job.cancel()

	return nil
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
		).
		First(ctx)

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

	if m.lw, err = worker.NewWorker(m, m.job.app.ID); err != nil {
		return err
	}

	if err = m.lw.Load(m.job.plugin); err != nil {
		return fmt.Errorf("failed load plugin: %v", err)
	}

	if err = m.lw.Run(ctx); err != nil {
		return err
	}

	return nil
}
