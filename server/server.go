package server

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/gobench-io/gobench/agent"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/worker"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	mu         sync.Mutex
	serverType serverType
	status     status
	master     master
	worker     worker.Worker

	isSchedule bool

	logger logger.Logger
}

// NewServer return a new server with provided options
func NewServer(opts *Options) (s *Server, err error) {
	// default db name
	dbFilename := "./gobench.sqlite3"

	s = &Server{
		serverType: opts.ServerType,
		isSchedule: true,
		logger:     opts.Logger,
	}

	if opts.ServerType == mtType {
		s.master.addr = opts.Addr
		s.master.port = opts.Port
		s.master.clusterPort = opts.ClusterPort
		s.master.dbFilename = dbFilename
		s.master.logger = s.logger
		s.master.la, err = agent.NewAgent(&s.master)
	}

	if opts.ServerType == wkType {
	}

	return s, nil
}

// Start begin a gobench server
func (s *Server) Start() (err error) {
	if err = s.master.setupDb(); err != nil {
		return
	}

	s.handleSignals()

	go s.master.schedule()

	// start the local agent socket server that communicate with local executor
	agentSocket := "/tmp/gobench-agentsocket"
	err = s.master.la.StartSocketServer(agentSocket)

	return
}

// DB returns the db client
func (s *Server) DB() *ent.Client {
	return s.master.db
}

func (s *Server) finish(status status) error {
	s.logger.Infow("server is shutting down")

	s.mu.Lock()
	s.status = status
	s.mu.Unlock()

	// todo: if there is a running scenario, shutdown
	// todo: send email if needed
	return s.master.db.Close()
}

// WebPort returns the master HTTP web port
func (s *Server) WebPort() int {
	return s.master.port
}

// NewApplication create a new application with a name and a scenario
// return the application id and error
func (s *Server) NewApplication(ctx context.Context, name, scenario string) (*ent.Application, error) {
	return s.master.db.Application.
		Create().
		SetName(name).
		SetScenario(scenario).
		SetStatus(string(jobPending)).
		Save(ctx)
}

// CancelApplication terminates an application
// if the app is running, send cancel signal
// if the app is finished/error, return ErrAppIsFinished error
// if the app is canceled, return with current app status
// else update app status with cancel
func (s *Server) CancelApplication(ctx context.Context, appID int) (*ent.Application, error) {
	err := s.master.cancel(ctx, appID)

	if err == nil {
		return s.master.db.Application.
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
	app, err := s.master.db.Application.
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
	return s.master.db.Application.
		UpdateOneID(appID).
		SetStatus(string(jobCancel)).
		Save(ctx)
}

// cleanupDB is the helper function to cleanup the DB for testing
func (s *Server) cleanupDB() error {
	ctx := context.TODO()
	_, err := s.master.db.Application.Delete().Exec(ctx)
	return err
}

// PrintAndDie print message to Stderr and exit error
func PrintAndDie(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// PrintVersionAndExit will print our version and exit.
func PrintVersionAndExit() {
	fmt.Printf("gobench: v%s\n", VERSION)
	os.Exit(0)
}
