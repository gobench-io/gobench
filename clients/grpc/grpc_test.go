package grpc

import (
	"errors"
	"testing"

	"github.com/gobench-io/gobench/executor"
	"github.com/gobench-io/gobench/executor/metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientConnect struct {
	mock.Mock
}

func (m *MockClientConnect) Setup(groups []metrics.Group) error {
	args := m.Called(groups)
	return args.Error(0)
}

func (m *MockClientConnect) Notify(title string, value int64) error {
	args := m.Called(title, value)
	return args.Error(0)
}

func TestGbClientConnSetupMethod(t *testing.T) {
	testClientConnect := new(MockClientConnect)
	executor.SetClientConnect(testClientConnect)

	expectedGroupsArg := []metrics.Group{
		{
			Name: "gRPC (host.io)",
			Graphs: []metrics.Graph{
				{
					Title: "gRPC Response",
					Unit:  "N",
					Metrics: []metrics.Metric{
						{
							Title: "get.foo.grpc_ok", // success
							Type:  metrics.Counter,
						},
						{
							Title: "get.foo.grpc_fail", // fail
							Type:  metrics.Counter,
						},
					},
				},
				{
					Title: "Latency",
					Unit:  "Microsecond",
					Metrics: []metrics.Metric{
						{
							Title: "get.foo.latency", // latency
							Type:  metrics.Histogram,
						},
					},
				},
			},
		},
	}

	testClientConnect.On("Setup", expectedGroupsArg).Return(nil)

	conn := &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          "host.io",
	}

	actualGraphs, err := conn.setupMethod("get.foo")
	assert.Nil(t, err)
	assert.Equal(t, expectedGroupsArg[0].Graphs, actualGraphs)
}

func TestGbClientConnSetupMethodError(t *testing.T) {
	testClientConnect := new(MockClientConnect)
	executor.SetClientConnect(testClientConnect)

	testClientConnect.On("Setup", mock.Anything).Return(errors.New("timeout"))

	conn := &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          "host.io",
	}

	_, err := conn.setupMethod("get.foo")
	assert.EqualError(t, err, "timeout")
}
