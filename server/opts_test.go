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
		Master:      true,
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
		isWorker:    true,
		Addr:        DEFAULT_HOST,
		Port:        0,
		Master:      false,
		ClusterPort: 0,
		Route:       "0.0.0.0:6890",
	}
	opts := &Options{
		isWorker: true,
	}
	setBaselineOptions(opts)
	assert.Equal(t, golden, opts)

}

func TestConfigureOptions(t *testing.T) {
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
}
