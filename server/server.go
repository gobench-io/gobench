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

type app struct {
	id       int
	name     string
	scenario string
}

type master struct {
	addr        string // host name
	port        int    // api port
	clusterPort int    // cluster port

	// database
	dbFilename string
	db         *ent.Client
}

type Server struct {
	mu         sync.Mutex
	serverType serverType
	status     status
	master     master
	worker     worker.Worker

	pendings []app // pending apps
	curr     *app  // current processing app

}

// NewServer return a new server with provided options
func NewServer(opts *Options) (*Server, error) {
	// default db name
	dbFilename := "./gobench.sqlite3"

	s := &Server{
		serverType: opts.ServerType,
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
	if err := s.setupDb(s.master.dbFilename); err != nil {
		return err
	}

	s.handleSignals()

	return nil
}

// NewApp creates a new application
// provided name and scenario
// returns application id and error
// func (s *Server) NewApp(name string, scenario string) (int, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	if s.curr == nil {
// 	}
// }

func (s *Server) setupDb(filename string) error {
	client, err := ent.Open(
		"sqlite3",
		filename+"?mode=rwc&cache=shared&&_busy_timeout=9999999&_fk=1")

	if err != nil {
		return fmt.Errorf("failed opening sqlite3 connection: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	s.mu.Lock()
	s.master.dbFilename = filename
	s.master.db = client
	s.mu.Unlock()

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
// func (s *Server) NewApplication(name, scenario string) (int, err) {
// }

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
