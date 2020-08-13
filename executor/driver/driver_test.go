package driver

import (
	"context"
	"testing"
	"time"

	"github.com/gobench-io/gobench/logger"
	"github.com/stretchr/testify/assert"
)

func loadValidPlugin(d *Driver) error {
	so := "./script/valid-dnt.so"
	return d.load(so)
}

func TestNew(t *testing.T) {
	so := "./script/valid-dnt/valid-dnt.so"

	n1, err := NewDriver(newNopMetricLog(), logger.NewNopLogger(), so, 1)
	assert.Nil(t, err)
	n2, err := NewDriver(newNopMetricLog(), logger.NewNopLogger(), so, 2)
	assert.Nil(t, err)

	assert.Equal(t, n1, n2)

	assert.Equal(t, n1.status, Idle)
	assert.Equal(t, n1.units, make(map[string]unit))
	assert.Len(t, *n1.vus, 1)

	assert.False(t, n1.Running())
}

func TestRunPlugin(t *testing.T) {
	so := "./script/valid-dnt/valid-dnt.so"
	n, _ := NewDriver(newNopMetricLog(), logger.NewNopLogger(), so, 1)

	assert.False(t, n.Running())

	ctx := context.Background()
	assert.Nil(t, n.Run(ctx))

	// after Run finish, the worker is in normal state
	assert.False(t, n.Running())
}

func TestCancelPlugin(t *testing.T) {
	so := "./script/valid-forever/valid-forever.so"
	n, _ := NewDriver(newNopMetricLog(), logger.NewNopLogger(), so, 1)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{}, 1)

	go func() {
		err := n.Run(ctx)

		assert.EqualError(t, err, ErrAppCancel.Error())
		assert.False(t, n.Running())

		done <- struct{}{}
	}()

	cancel()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("Should have finish the running after cancel")
	}
}
