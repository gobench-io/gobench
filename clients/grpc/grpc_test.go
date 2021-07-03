package gbGrpc

import (
	"testing"

	"github.com/gobench-io/gobench/executor/metrics"
	"github.com/stretchr/testify/assert"
)

// TODO: need to mock executorInstance
// executor.SetInstance(executor interface)
func TestGbClientConnSetupMethod(t *testing.T) {
	conn := &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          "host.io",
	}

	method := "get.foo"
	graphs, err := conn.setupMethod(method)
	assert.Nil(t, err)
	assert.Equal(t, graphs, []metrics.Graph{
		{
			Title: "gRPC Response",
			Unit:  "N",
			Metrics: []metrics.Metric{
				{
					Title: method + ".grpc_ok", // success
					Type:  metrics.Counter,
				},
				{
					Title: method + ".grpc_fail", // fail
					Type:  metrics.Counter,
				},
			},
		},
	})
}
