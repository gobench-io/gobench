package main

import (
	"flag"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testOpts struct {
	args          []string
	version, help func()
}

func TestVersionHelp(t *testing.T) {
	ch := make(chan bool, 1)

	checkPrintInvoked := func() {
		ch <- true
	}

	usage := func() {
		panic("should not get there")
	}

	// test the help
	testHelps := []testOpts{
		{[]string{"--version"}, checkPrintInvoked, usage},
		{[]string{"-v"}, checkPrintInvoked, usage},
		{[]string{"--help"}, usage, checkPrintInvoked},
		{[]string{"-h"}, usage, checkPrintInvoked},
	}

	for _, tt := range testHelps {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		opts, err := ConfigureOptions(fs, tt.args, tt.version, tt.help)
		assert.Nil(t, err)
		assert.Nil(t, opts)

		select {
		case <-ch:
		case <-time.After(time.Second):
			assert.Fail(t, "should have invoked print function for args %s", tt.args)
		}
	}
}

func TestExecutorOption(t *testing.T) {
	version := func() {}
	help := func() {}

	testFuncs := []testOpts{
		{[]string{
			"--mode", "executor",
			"--agent-sock", "agent/sock",
			"--executor-sock", "executor/sock",
			"--driver-path", "driver/path",
			"--app-id", "123",
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
			AppID:        123,
		}, *opts)
	}

	// missing parameter tests
	testErrs := []testOpts{
		// {[]string{
		// 	"--agent-sock", "agent/sock",
		// 	"--executor-sock", "executor/sock",
		// 	"--driver-path", "driver/path",
		// 	"--app-id", "123",
		// }, version, help},
		{[]string{
			"--mode", "executor",
			"--executor-sock", "executor/sock",
			"--driver-path", "driver/path",
		}, version, help},
		{[]string{
			"--mode", "executor",
			"--agent-sock", "agent/sock",
			"--driver-path", "driver/path",
		}, version, help},
		{[]string{
			"--mode", "executor",
			"--agent-sock", "agent/sock",
			"--executor-sock", "executor/sock",
		}, version, help},
	}

	for _, tt := range testErrs {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		opts, err := ConfigureOptions(fs, tt.args, tt.version, tt.help)
		assert.EqualError(t, err, ErrInvalidFlags.Error())
		assert.Nil(t, opts)
	}
}

func TestMasterOption(t *testing.T) {
	// helper function
	mustNotFail := func(args []string) *Options {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		opts, err := ConfigureOptions(fs, args, usage, usage)
		if err != nil {
			t.Fatalf("Error on config: %v", err)
		}
		return opts
	}

	mustFail := func(args []string, errContent ...string) {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		opts, err := ConfigureOptions(fs, args, usage, usage)
		if opts != nil || err == nil {
			t.Fatalf("Expect no opts and err, got %v and %v", opts, err)
		}
		for _, content := range errContent {
			if strings.Contains(err.Error(), content) {
				return
			}
		}
		t.Fatalf("Expect error contain any of %v, got %v", errContent, err)
	}

	// check mode
	mustFail([]string{"--mode", "not existed"}, "mode must be either master, agent, or executor")

	opts := mustNotFail([]string{"-p", "3000"})
	assert.Equal(t, opts.Port, 3000)
}
