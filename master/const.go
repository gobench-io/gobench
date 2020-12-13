package master

import "errors"

type status string

const (
	statusInit    status = "init"
	statusRunning status = "running"
	statusIdle    status = "idle"
	statusCancel  status = "cancel"
)

// App states
const (
	jobPending      jobState = "pending"
	jobProvisioning jobState = "provisioning"
	jobRunning      jobState = "running"
	jobFinished     jobState = "finished"
	jobCancel       jobState = "cancel"
	jobError        jobState = "error"
)

// Error
var (
	ErrAppNotRunning = errors.New("application is not running")
	ErrAppIsFinished = errors.New("application is finished already")
	ErrAppIsCanceled = errors.New("application is canceled")
	ErrCantDeleteApp = errors.New("cannot delete a %s application")
)

var (
	// gitCommit injected at build
	gitCommit string
)

const (
	// VERSION is the current version for the server.
	VERSION = "0.0.1"
)
