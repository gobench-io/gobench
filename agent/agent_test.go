package agent

import (
	"testing"

	"github.com/gobench-io/gobench/logger"
	"github.com/stretchr/testify/assert"
)

func TestNewAgent(t *testing.T) {
	opts := &Options{
		Route:       "localhost:1234",
		ClusterPort: 2345,
	}
	logger := logger.NewNopLogger()
	ml := newNopMetricLog()

	_, err := NewAgent(opts, ml, logger)
	assert.Nil(t, err)
}

func TestStartAgent(t *testing.T) {
	// new agent
	// start socket server
	// start http server
}
