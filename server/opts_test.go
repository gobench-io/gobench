package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultOptions(t *testing.T) {
	golden := &Options{
		Addr:        DEFAULT_HOST,
		Port:        DEFAULT_PORT,
		Master:      true,
		ClusterPort: DEFAULT_CLUSTER_PORT,
		Route:       "",
	}

	opts := &Options{}
	setBaselineOptions(opts)
	assert.Equal(t, golden, opts)

	// worker role
	golden = &Options{
		isWorker:    true,
		Addr:        DEFAULT_HOST,
		Port:        0,
		Master:      false,
		ClusterPort: 0,
		Route:       "0.0.0.0:6890",
	}
	opts = &Options{
		isWorker: true,
	}
	setBaselineOptions(opts)
	assert.Equal(t, golden, opts)
}
