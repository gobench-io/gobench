package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gobench-io/gobench/ent"
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
}

// NewServer return a new server with provided options
func NewServer(opts *Options) (*Server, error) {
	// default db name
	dbFilename := "./gobench.sqlite3"

	s := &Server{
		serverType: opts.ServerType,
		isSchedule: true,
	}

	if opts.ServerType == mtType {
		s.master.addr = opts.Addr
		s.master.port = opts.Port
		s.master.clusterPort = opts.ClusterPort
		s.master.dbFilename = dbFilename
	}

	if opts.ServerType == wkType {
	}

	return s, nil
}

// Start begin a gobench server
func (s *Server) Start() error {
	if err := s.master.setupDb(); err != nil {
		return err
	}

	s.handleSignals()

	go s.master.schedule()

	return nil
}

// DB returns the db client
func (s *Server) DB() *ent.Client {
	return s.master.db
}

func (s *Server) finish(status status) error {
	log.Println("server is shutting down")

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
