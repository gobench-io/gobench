package executor

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	er, _ := newExecutorRPC(e)

	args := true
	reply := new(bool)

	err = er.Start(&args, reply)
	assert.Nil(t, err)
}

func TestGenerate(t *testing.T) {
	dir, err := ioutil.TempDir("", "scenario-*")
	assert.Nil(t, err)
	name := filepath.Join(dir, "main.go")
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	assert.Nil(t, err)

	err = Generate(f, "agentsocket", "executorsocket", 123)
	assert.Nil(t, err)
	log.Println("dir", dir)
	os.Remove(name)
}

// a generated file should be compiled with a valid scenario
func TestCompile(t *testing.T) {

}
