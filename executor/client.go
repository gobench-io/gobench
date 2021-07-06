package executor

import (
	"github.com/gobench-io/gobench/executor/metrics"
)

type ClientConnector interface {
	Setup(groups []metrics.Group) error
	Notify(title string, value int64) error
}

// Setup is used for the driver to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	clientConnect := getClientConnect()

	return clientConnect.Setup(groups)
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	clientConnect := getClientConnect()

	return clientConnect.Notify(title, value)
}
