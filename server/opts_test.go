package server

import (
	"flag"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultOptions(t *testing.T) {
	golden := &Options{
		Addr:        DEFAULT_HOST,
		Port:        DEFAULT_PORT,
		ServerType:  master,
		ClusterPort: DEFAULT_CLUSTER_PORT,
		Route:       "",
	}

	opts := &Options{}
	setBaselineOptions(opts)
	assert.Equal(t, golden, opts)
}

func TestDefaultWorkerOptions(t *testing.T) {
	// worker role
	golden := &Options{
		Addr:        DEFAULT_HOST,
		Port:        0,
		ServerType:  worker,
		ClusterPort: 0,
		Route:       "0.0.0.0:6890",
	}
	opts := &Options{
		ServerType: worker,
	}
	setBaselineOptions(opts)
	assert.Equal(t, golden, opts)

}

func TestConfigureOptions(t *testing.T) {
	// helper function
	ch := make(chan bool, 1)
	checkPrintInvoked := func() {
		ch <- true
	}

	usage := func() {
		panic("should not get there")
	}

	type testPrint struct {
		args          []string
		version, help func()
	}

	testFuncs := []testPrint{
		{[]string{"-v"}, checkPrintInvoked, usage},
		{[]string{"--version"}, checkPrintInvoked, usage},
		{[]string{"-h"}, usage, checkPrintInvoked},
		{[]string{"--help"}, usage, checkPrintInvoked},
	}

	for _, tf := range testFuncs {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		opts, err := ConfigureOptions(fs, tf.args, tf.version, tf.help)
		if err != nil {
			t.Fatalf("Error on config: %v", err)
		}
		if opts != nil {
			t.Fatalf("Expected options to be nil, got %v for args %s", opts, tf.args)
		}
		select {
		case <-ch:
		case <-time.After(1 * time.Second):
			t.Fatalf("Should have invoked print function for args=%v", tf.args)
		}
	}

	// helper function
	mustNotFail := func(args []string) *Options {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		opts, err := ConfigureOptions(fs, args, usage, usage)
		if err != nil {
			t.Fatalf("Error on config: %v", err)
		}
		return opts
	}

	opts := mustNotFail([]string{"-p", "3000"})
	assert.Equal(t, opts.Port, 3000)

	opts = mustNotFail([]string{"-m", "true"})
	assert.Equal(t, opts.ServerType, master)
	opts = mustNotFail([]string{"-m"})
	assert.Equal(t, opts.ServerType, master)

	opts = mustNotFail([]string{"-w", "true"})
	assert.Equal(t, opts.ServerType, worker)
	opts = mustNotFail([]string{"-w"})
	assert.Equal(t, opts.ServerType, worker)
}
