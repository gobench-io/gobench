package executor

import (
	"fmt"
	"net"
	"os"

	"github.com/gobench-io/gobench/executor/driver"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/pb"
	"github.com/gobench-io/gobench/scenario"
	"google.golang.org/grpc"
)

// Options is for creating new executor object
type Options struct {
	AgentSock    string
	ExecutorSock string
	AppID        int
	Vus          scenario.Vus
}

// Executor struct
type Executor struct {
	id           string
	logger       logger.Logger
	agentSock    string
	executorSock string
	appID        int

	driver *driver.Driver
	rc     pb.AgentClient
}

// NewExecutor creates a new executor
// also load the plugin from driver path
func NewExecutor(opts *Options, logger logger.Logger) (e *Executor, err error) {
	// id is the combination of hostname and pid
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	id := fmt.Sprintf("%s-%d", hostname, pid)

	e = &Executor{
		id:           id,
		logger:       logger,
		agentSock:    opts.AgentSock,
		executorSock: opts.ExecutorSock,
		appID:        opts.AppID,
	}

	e.driver, err = driver.NewDriver(e, logger, opts.Vus, opts.AppID)

	return
}

// Serve starts a rpc server at the executor socket
// and connects to the agent via agent socket
func (e *Executor) Serve() (err error) {
	// establishes a connection to agent rpc server
	socket := "passthrough:///unix://" + e.agentSock
	conn, err := grpc.Dial(socket, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	e.rc = pb.NewAgentClient(conn)

	// executor register a rpc server at executor socket
	l, err := net.Listen("unix", e.executorSock)
	if err != nil {
		return
	}

	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, e)

	err = s.Serve(l)
	if err != nil {
		return
	}

	return
}
