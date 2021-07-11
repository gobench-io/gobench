package gbGrpc

import (
	"context"
	"errors"
	"testing"

	"github.com/gobench-io/gobench/executor"
	"github.com/gobench-io/gobench/executor/metrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type mockExecutor struct {
	mock.Mock
}

func (m *mockExecutor) Setup(groups []metrics.Group) error {
	args := m.Called(groups)
	return args.Error(0)
}

func (m *mockExecutor) Notify(title string, value int64) error {
	args := m.Called(title, value)
	return args.Error(0)
}

type mockGbClientStream struct {
	mock.Mock
}

// func (m *mockGbClientStream) Setup(groups []metrics.Group) error {
// 	args := m.Called(groups)
// 	return args.Error(0)
// }

// func (m *mockGbClientStream) Notify(title string, value int64) error {
// 	args := m.Called(title, value)
// 	return args.Error(0)
// }

type mockClientConnect struct {
	mock.Mock
}

func (m *mockClientConnect) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	rets := m.Called(ctx, method, args, reply, opts)
	return rets.Error(0)
}

func (m *mockClientConnect) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	rets := m.Called(ctx, desc, method, opts)
	return rets.Get(0).(grpc.ClientStream), rets.Error(1)
}

func TestGbClientConnSetupMethod(t *testing.T) {
	me := new(mockExecutor)
	executor.SetClientConnect(me)

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

	me.On("Setup", expectedGroupsArg).Return(nil)

	conn := &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          "host.io",
	}

	actualGraphs, err := conn.setupMethod("get.foo")
	assert.Nil(t, err)
	assert.Equal(t, expectedGroupsArg[0].Graphs, actualGraphs)
}

func TestGbClientConnSetupMethodError(t *testing.T) {
	me := new(mockExecutor)
	executor.SetClientConnect(me)

	me.On("Setup", mock.Anything).Return(errors.New("timeout"))

	conn := &GbClientConn{
		methodGraphsMap: make(map[string][]metrics.Graph),
		target:          "host.io",
	}

	_, err := conn.setupMethod("get.foo")
	assert.EqualError(t, err, "timeout")
}

// func TestGbClientConnInvok(t *testing.T) {
// 	me := new(mockExecutor)
// 	executor.SetClientConnect(me)

// 	me.On("Setup", mock.Anything).Return(nil)
// 	me.On("Notify", "get.foo.latency", mock.Anything).Return(nil)
// 	me.On("Notify", "get.foo.grpc_ok", 1).Return(nil)

// 	conn := &GbClientConn{
// 		methodGraphsMap: make(map[string][]metrics.Graph),
// 		target:          "host.io",
// 	}

// 	_, err := conn.setupMethod("get.foo")
// 	assert.Nil(t, err)

// 	ctx := context.Background()
// 	var in, out interface{}

// 	err = conn.Invoke(ctx, "get.foo", in, out)
// 	assert.Nil(t, err)
// }

func TestGbClientStreamSetupMethod(t *testing.T) {
	me := new(mockExecutor)
	executor.SetClientConnect(me)

	expectedGroupsArg := []metrics.Group{
		{
			Name: "gRPC Stream (host.io)",
			Graphs: []metrics.Graph{
				{
					Title: "New Stream",
					Unit:  "N",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.new_stream_ok",
							Type:  metrics.Counter,
						},
						{
							Title: "list.foo.new_stream_fail",
							Type:  metrics.Counter,
						},
					},
				},
				{
					Title: "New Stream Latency",
					Unit:  "Microsecond",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.new_stream_latency",
							Type:  metrics.Histogram,
						},
					},
				},
				{
					Title: "Send Message",
					Unit:  "N",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.send_msg_ok",
							Type:  metrics.Counter,
						},
						{
							Title: "list.foo.send_msg_fail",
							Type:  metrics.Counter,
						},
					},
				},
				{
					Title: "Send Message Latency",
					Unit:  "Microsecond",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.send_msg_latency",
							Type:  metrics.Histogram,
						},
					},
				},
				{
					Title: "Receive Message",
					Unit:  "N",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.recv_msg_ok",
							Type:  metrics.Counter,
						},
						{
							Title: "list.foo.recv_msg_fail",
							Type:  metrics.Counter,
						},
					},
				},
				{
					Title: "Receive Message Latency",
					Unit:  "Microsecond",
					Metrics: []metrics.Metric{
						{
							Title: "list.foo.recv_msg_latency",
							Type:  metrics.Histogram,
						},
					},
				},
			},
		},
	}

	me.On("Setup", expectedGroupsArg).Return(nil)

	cs := &GbClientStream{}

	actualGraphs, err := cs.setupMethod("host.io", "list.foo")
	assert.Nil(t, err)
	assert.Equal(t, expectedGroupsArg[0].Graphs, actualGraphs)
}
