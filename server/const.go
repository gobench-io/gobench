package server

import "errors"

const (
	// VERSION is the current version for the server.
	VERSION = "0.0.1"

	DEFAULT_HOST         = "0.0.0.0"
	DEFAULT_PORT         = 8080
	DEFAULT_CLUSTER_PORT = 8081
)

type status string

const (
	statusInit    status = "init"
	statusRunning status = "running"
	statusIdle    status = "idle"
	statusCancel  status = "cancel"
)

// Error
var (
	ErrAppNotRunning = errors.New("application is not running")
	ErrAppIsFinished = errors.New("application is finished already")
	ErrCantDeleteApp = errors.New("cannot delete an application is %s")
)
