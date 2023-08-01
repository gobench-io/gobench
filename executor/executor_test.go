package executor

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/gobench-io/gobench/v2/executor/scenario"
	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/pb"
	"github.com/stretchr/testify/assert"
)

func generate(t *testing.T) (string, string) {
	dir, err := ioutil.TempDir("", "scenario-*")
	assert.Nil(t, err)
	name := filepath.Join(dir, "main.go")
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	assert.Nil(t, err)

	err = Generate(f, 123)
	assert.Nil(t, err)

	return dir, name
}

func TestGenerate(t *testing.T) {
	dir, _ := generate(t)
	os.RemoveAll(dir)
}

func TestNew(t *testing.T) {
	vus := scenario.Vus{
		scenario.Vu{

			Nu:   20,
			Rate: 100,
			Fu:   func(ctx context.Context, vui int) {},
		},
	}

	e1, err := NewExecutor(&Options{Vus: vus}, logger.NewNopLogger())
	assert.Nil(t, err)
	e2, err := NewExecutor(&Options{Vus: vus}, logger.NewNopLogger())
	assert.Nil(t, err)

	// singleton object
	assert.Equal(t, e1, e2)

	assert.Equal(t, e1.status, Idle)
	assert.Equal(t, e1.units, make(map[string]unit))
	assert.Len(t, e1.vus, 1)
}

// a generated file should be compiled with a valid scenario
func TestCompile(t *testing.T) {
	scenario := `
package main

import (
	"context"

	"github.com/gobench-io/gobench/v2/executor/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   1,
			Rate: 100,
			Fu:   f1,
		},
	}
}

func f1(ctx context.Context, vui int) {
}`

	dir, _ := generate(t)

	defer os.RemoveAll(dir)

	// create scenario.go
	scenarioPath := filepath.Join(dir, "scenario.go")
	scenarioFile, _ := os.OpenFile(scenarioPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	_, _ = scenarioFile.Write([]byte(scenario))

	// create go.mod
	testDir, _ := os.Getwd()
	mainDir, _ := exec.Command("dirname", testDir).CombinedOutput()
	gomod := fmt.Sprintf(`
		module gobench.io/scenario
		replace github.com/gobench-io/gobench/v2 => %s
		`, string(mainDir))
	gomodPath := filepath.Join(dir, "go.mod")
	gomodFile, _ := os.OpenFile(gomodPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	_, _ = gomodFile.Write([]byte(gomod))

	out, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s; go get; go build -o main.out", dir)).CombinedOutput()
	assert.Nil(t, err, string(out))
}

func TestStart(t *testing.T) {
	opts := &Options{
		AgentSock:    "/tmp/a1",
		ExecutorSock: "/tmp/e1",
		AppID:        1,
		Vus: scenario.Vus{
			scenario.Vu{
				Nu:   20,
				Rate: 100,
				Fu:   func(ctx context.Context, vui int) {},
			},
		},
	}
	logger := logger.NewNopLogger()

	e, err := NewExecutor(opts, logger)
	assert.Nil(t, err)

	e.rc = newNopMetricLog()

	ctx := context.TODO()

	_, err = e.Start(ctx, &pb.StartRequest{
		AppID: int64(opts.AppID),
	})

	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	opts := &Options{
		AgentSock:    "/tmp/a1",
		ExecutorSock: "/tmp/e1",
		AppID:        1,
		Vus: scenario.Vus{
			scenario.Vu{
				Nu:   20,
				Rate: 100,
				Fu: func(ctx context.Context, vui int) {
					for {
						time.Sleep(time.Second)
					}
				},
			},
		},
	}
	logger := logger.NewNopLogger()

	e, err := NewExecutor(opts, logger)
	assert.Nil(t, err)

	e.rc = newNopMetricLog()

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{}, 1)

	go func() {
		_, err = e.Start(ctx, &pb.StartRequest{
			AppID: int64(opts.AppID),
		})
		assert.EqualError(t, err, ErrAppCancel.Error())
		assert.Equal(t, Finished, e.status)

		done <- struct{}{}
	}()

	cancel()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatalf("Should have finish the running after cancel")
	}
}
