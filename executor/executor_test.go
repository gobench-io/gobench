package executor

import (
	"testing"

	"github.com/gobench-io/gobench/logger"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	opts := &Options{
		AgentSock:    "/tmp/a1",
		ExecutorSock: "/tmp/e1",
		DriverPath:   "./driver/script/valid-dnt/valid-dnt.so",
		AppID:        1,
	}
	logger := logger.NewNopLogger()

	e, err := NewExecutor(opts, logger)
	assert.Nil(t, err)

	// setup nop metric logger for the driver
	assert.Nil(t, e.driver.SetNopMetricLog())

	args := true
	reply := new(bool)

	err = e.Start(&args, reply)
	assert.Nil(t, err)
}
