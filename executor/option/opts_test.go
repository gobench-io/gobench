package option

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureOptions(t *testing.T) {
	type testPrint struct {
		args          []string
		version, help func()
	}

	version := func() {}
	help := func() {}

	testFuncs := []testPrint{
		{[]string{
			"--agent-sock", "agent/sock",
			"--executor-sock", "executor/sock",
			"--driver-path", "driver/path",
		}, version, help},
	}

	for _, tt := range testFuncs {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		opts, err := ConfigureOptions(fs, tt.args, tt.version, tt.help)
		assert.Nil(t, err, "err on ConfigureOptions")
		assert.NotNil(t, opts)
		assert.Equal(t, Options{
			AgentSock:    "agent/sock",
			ExecutorSock: "executor/sock",
			DriverPath:   "driver/path",
		}, *opts)
	}

	// missing parameter tests
	testErr := []testPrint{
		{[]string{
			"--executor-sock", "executor/sock",
			"--driver-path", "driver/path",
		}, version, help},
		{[]string{
			"--agent-sock", "agent/sock",
			"--driver-path", "driver/path",
		}, version, help},
		{[]string{
			"--agent-sock", "agent/sock",
			"--executor-sock", "executor/sock",
		}, version, help},
	}

	for _, tt := range testErr {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		opts, err := ConfigureOptions(fs, tt.args, tt.version, tt.help)
		assert.EqualError(t, err, ErrInvalidFlags.Error())
		assert.Nil(t, opts)
	}
}
