package agent

import (
	"testing"

	"github.com/gobench-io/gobench/v2/logger"
	"github.com/stretchr/testify/assert"
)

func newAgent(t *testing.T, opts *Options) *Agent {
	logger := logger.NewNopLogger()
	ml := newNopMetricLog()

	a, err := NewAgent(opts, ml, logger)
	assert.Nil(t, err)

	return a
}

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
	a := newAgent(t, &Options{
		Route:       "localhost:1234",
		ClusterPort: 2345,
	})
	assert.Nil(t, a.StartSocketServer())
	// insert the grpc over tcp here
}
