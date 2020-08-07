package executor

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/gobench-io/gobench/executor/driver"
	"github.com/gobench-io/gobench/logger"
)

type Options struct {
	AgentSock    string
	ExecutorSock string
	DriverPath   string // the plugin user wrote
	AppID        int
}

type Executor struct {
	logger       logger.Logger
	agentSock    string
	executorSock string
	appID        int

	driver *driver.Driver
	rc     *rpc.Client
}

func NewExecutor(opts *Options, logger logger.Logger) (e *Executor, err error) {
	e = &Executor{
		logger:       logger,
		agentSock:    opts.AgentSock,
		executorSock: opts.ExecutorSock,
		appID:        opts.AppID,
	}

	e.driver, err = driver.NewDriver(e, logger, opts.DriverPath, opts.AppID)

	return
}

// Serve starts a rpc server at the executor socket
// and connect to agent via agent socket
func (e *Executor) Serve() (err error) {
	// establishes a connection
	e.rc, err = rpc.DialHTTP("unix", e.agentSock)
	if err != nil {
		log.Println(err)
		return
	}

	// register a rpc server at executor socket
	err = rpc.Register(e)
	if err != nil {
		return
	}
	rpc.HandleHTTP()

	// bind rpc to executor sock
	os.Remove(e.executorSock)
	l, err := net.Listen("unix", e.executorSock)
	if err != nil {
		return
	}
	err = http.Serve(l, nil)

	return
}
