package master

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"context"

	"github.com/gobench-io/gobench/agent"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/logger"

	_ "github.com/mattn/go-sqlite3"
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

type Master struct {
	mu          sync.Mutex
	addr        string // host name
	port        int    // api port
	clusterPort int    // cluster port

	status  status
	logger  logger.Logger
	program string

	// database
	isScheduled bool
	dbFilename  string
	db          *ent.Client

	la  *agent.Agent // local agent
	job *job
}

type job struct {
	app    *ent.Application
	plugin string // plugin path
	cancel context.CancelFunc
}

type Options struct {
	Port    int
	Addr    string
	DbPath  string
	Program string
}

func NewMaster(opts *Options, logger logger.Logger) (m *Master, err error) {
	logger.Infow("new master program",
		"port", opts.Port,
		"db file path", opts.DbPath,
	)

	m = &Master{
		addr:       opts.Addr,
		port:       opts.Port,
		dbFilename: opts.DbPath,
		logger:     logger,
		program:    opts.Program,
	}

	m.isScheduled = true // by default

	agentSocket := fmt.Sprintf("/tmp/gobench-agentsocket-%d", os.Getpid())
	la, err := agent.NewAgent(&agent.Options{Socket: agentSocket}, m, logger)
	if err != nil {
		return
	}
	m.la = la

	return
}

func (m *Master) Start() (err error) {
	if err = m.setupDb(); err != nil {
		return
	}

	m.handleSignals()

	if m.isScheduled {
		go m.schedule()
	}

	// create a go.sum if not existed to avoid "go: cannot find main module, but
	// -modfile was set" error
	cmd := exec.Command("touch", "go.mod")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		m.logger.Errorw("touch go.mod",
			"stdoutStderr", stdoutStderr,
			"err", err)
		return err
	}

	// start the local agent socket server that communicate with local executor
	err = m.la.StartSocketServer()

	return
}

// DB returns the db client
func (m *Master) DB() *ent.Client {
	return m.db
}

func (m *Master) finish(status status) error {
	m.logger.Infow("server is shutting down")

	m.mu.Lock()
	m.status = status
	m.mu.Unlock()

	// todo: if there is a running scenario, shutdown
	// todo: send email if needed
	return m.db.Close()
}

// WebPort returns the master HTTP web port
func (m *Master) WebPort() int {
	return m.port
}

// NewApplication create a new application with a name and a scenario
// return the application id and error
func (m *Master) NewApplication(ctx context.Context, name, scenario, gomod, gosum string) (
	*ent.Application, error,
) {
	return m.db.Application.
		Create().
		SetName(name).
		SetScenario(scenario).
		SetGomod(gomod).
		SetGosum(gosum).
		SetStatus(string(jobPending)).
		Save(ctx)
}

// SetApplicationTags set application tags
func (m *Master) SetApplicationTags(ctx context.Context, appID int, tags string) (*ent.Application, error) {
	return m.db.Application.
		UpdateOneID(appID).
		SetTags(tags).
		Save(ctx)
}

// DeleteApplication a pending/finished/canceled/error application
func (m *Master) DeleteApplication(ctx context.Context, appID int) error {
	app, err := m.db.Application.
		Query().
		Where(application.ID(appID)).
		Only(ctx)

	if err != nil {
		return err
	}

	if app.Status != string(jobPending) && app.Status != string(jobCancel) &&
		app.Status != string(jobFinished) && app.Status != string(jobError) {
		return fmt.Errorf(ErrCantDeleteApp.Error(), string(app.Status))
	}

	return m.db.Application.
		DeleteOneID(appID).
		Exec(ctx)
}

// CancelApplication terminates an application
// if the app is running, send cancel signal
// if the app is finished/error, return ErrAppIsFinished error
// if the app is canceled, return with current app status
// else update app status with cancel
func (m *Master) CancelApplication(ctx context.Context, appID int) (*ent.Application, error) {
	err := m.cancel(ctx, appID)

	if err == nil {
		return m.db.Application.
			Query().
			Where(application.ID(appID)).
			Only(ctx)
	}

	// if err and err is not the app is not running
	if err != nil && !errors.Is(err, ErrAppNotRunning) {
		return nil, err
	}

	// if the app is not running, update directly on the db
	// query the app
	// if the app status is finished or error, return error
	// if the app status is cancel (already), just return
	// else, update the app table
	app, err := m.db.Application.
		Query().
		Where(application.ID(appID)).
		Only(ctx)

	if err != nil {
		return app, err
	}

	if app.Status == string(jobCancel) {
		return app, nil
	}
	if app.Status == string(jobFinished) || app.Status == string(jobError) {
		return app, ErrAppIsFinished
	}

	// else, update the status on db
	return m.db.Application.
		UpdateOneID(appID).
		SetStatus(string(jobCancel)).
		Save(ctx)
}

// cleanupDB is the helper function to cleanup the DB for testing
func (m *Master) cleanupDB() error {
	ctx := context.TODO()
	_, err := m.db.Application.Delete().Exec(ctx)
	return err
}

// to is the function to set new state for an application
// save new state to the db
func (m *Master) jobTo(ctx context.Context, state jobState) (err error) {
	m.job.app, err = m.job.app.Update().
		SetStatus(string(state)).
		Save(ctx)

	return
}

// setupDb setup the db in the master
func (m *Master) setupDb() error {
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

	m.db = client

	return nil
}

// schedule get a pending application from the db if there is no active job
func (m *Master) schedule() {
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

func (m *Master) run(ctx context.Context, j *job) (err error) {
	// create new job from the application
	m.job = j

	defer func() {
		je := jobFinished

		// normalize je
		if err != nil {
			m.logger.Infow("failed run job",
				"application id", m.job.app.ID,
				"err", err,
			)
			je = jobError

			if ctx.Err() != nil {
				je = jobCancel
				err = ErrAppIsCanceled
			}
		}

		// create new context
		ctx := context.TODO()
		_ = m.jobTo(ctx, je)

		m.logger.Infow("job new status",
			"application id", m.job.app.ID,
			"status", m.job.app.Status,
		)
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

	return
}

// cancel terminates a running job with the same app ID
func (m *Master) cancel(ctx context.Context, appID int) error {
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
func (m *Master) provision() (*ent.Application, error) {
	// compile
	return nil, nil
}

func (m *Master) nextApplication(ctx context.Context) (*ent.Application, error) {
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

func saveToFile(content []byte, dir, file string) (name string, err error) {
	// save the scenario to a tmp file
	name = filepath.Join(dir, file)
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return
	}

	if _, err = f.Write(content); err != nil {
		return
	}

	err = f.Close()

	return
}

// jobCompile using go to compile a scenario in plugin build mode
// the result is path to so file
func (m *Master) jobCompile(ctx context.Context) error {
	var path string

	scen := m.job.app.Scenario
	gomod := m.job.app.Gomod
	gosum := m.job.app.Gosum

	dir, err := ioutil.TempDir("", "scenario-*")
	if err != nil {
		return fmt.Errorf("create temp dir: %v", err)
	}

	m.logger.Infow("folder for compiling",
		"dir", dir)

	// save the scenario to a tmp file
	tmpScenName, err := saveToFile([]byte(scen), dir, "scenario.go")
	if err != nil {
		return err
	}
	defer os.Remove(tmpScenName) // cleanup

	var tmpGomodName string
	if gomod != "" {
		tmpGomodName, err = saveToFile([]byte(gomod), dir, "go.mod")
		if err != nil {
			return err
		}
		defer os.Remove(tmpGomodName) // cleanup
	}
	var tmpGosumName string
	if gosum != "" {
		if tmpGosumName, err = saveToFile([]byte(gosum), dir, "go.sum"); err != nil {
			return err
		}
		defer os.Remove(tmpGosumName)
	}

	path = fmt.Sprintf("%s.out", tmpScenName)

	// compile the scenario to a tmp file
	exe := "go"
	args := []string{
		"build",
		"-buildmode=plugin",
	}
	if gomod != "" {
		args = append(args, "-modfile="+tmpGomodName)
	}
	args = append(args, "-o", path, tmpScenName)

	cmd := exec.Command(exe, args...)

	if out, err := cmd.CombinedOutput(); err != nil {
		m.logger.Errorw("failed compiling the scenario",
			"err", err,
			"output", string(out))
		return fmt.Errorf("compile scenario: %v", err)
	}

	m.job.plugin = path

	return nil
}

// runJob runs the already compiled plugin, uses agent workhouse
func (m *Master) runJob(ctx context.Context) (err error) {
	return m.la.RunJob(ctx, m.program, m.job.plugin, m.job.app.ID)
}
