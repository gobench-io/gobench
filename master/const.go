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
	// gitCommit, gitTag injected at build
	gitCommit string
	gitTag    string
)
