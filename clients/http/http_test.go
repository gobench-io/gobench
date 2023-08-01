package http

import (
	"context"
	"errors"
	"testing"

	"github.com/gobench-io/gobench/v2/executor"
	"github.com/gobench-io/gobench/v2/executor/metrics"
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

func TestNewHttpClient(t *testing.T) {
	testClientConnect := new(MockClientConnect)
	executor.SetClientConnect(testClientConnect)

	expectedGroupsArg := []metrics.Group{
		{
			Name: "HTTP (foo)",
			Graphs: []metrics.Graph{
				{
					Title: "HTTP Response",
					Unit:  "N",
					Metrics: []metrics.Metric{
						{
							Title: "foo.http_ok",
							Type:  metrics.Counter,
						},
						{
							Title: "foo.http_fail",
							Type:  metrics.Counter,
						},
						{
							Title: "foo.http_other_fail",
							Type:  metrics.Counter,
						},
					},
				},
				{
					Title: "Latency",
					Unit:  "Microsecond",
					Metrics: []metrics.Metric{
						{
							Title: "foo.latency",
							Type:  metrics.Histogram,
						},
					},
				},
			},
		},
	}
	testClientConnect.On("Setup", expectedGroupsArg).Return(nil)

	ctx := context.Background()

	httpClient, err := NewHttpClient(ctx, "foo")

	testClientConnect.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, "foo", httpClient.prefix)
}

func TestNewHttpClientError(t *testing.T) {
	testClientConnect := new(MockClientConnect)
	executor.SetClientConnect(testClientConnect)

	testClientConnect.On("Setup", mock.Anything).Return(errors.New("timeout"))

	ctx := context.Background()

	_, err := NewHttpClient(ctx, "foo")

	testClientConnect.AssertExpectations(t)
	assert.EqualError(t, err, "timeout")
}
