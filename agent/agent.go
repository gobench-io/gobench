package agent

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
)

type metricLoggerRPC interface {
	FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error
	FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error
	FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error
	Counter(req *metrics.CounterReq, res *metrics.CounterRes) error
	Histogram(req *metrics.HistogramReq, res *metrics.HistogramRes) error
	Gauge(req *metrics.GaugeReq, res *metrics.GaugeRes) error
}

type Agent struct {
	// when this is the local agent, inherit from master
	// when this is the remote agent, ... todo
	metricLoggerRPC
	logger logger.Logger
	socket string
	rs     *rpc.Server
}

func NewAgent(ml metricLoggerRPC, logger logger.Logger) (*Agent, error) {
	a := &Agent{
		metricLoggerRPC: ml,
		logger:          logger,
		rs:              rpc.NewServer(),
	}
	return a, nil
}

func (a *Agent) StartSocketServer(socket string) error {
	a.socket = socket

	a.rs.Register(a)

	serverMux := http.NewServeMux()
	serverMux.Handle(rpc.DefaultRPCPath, a.rs)
	serverMux.Handle(rpc.DefaultDebugPath, a.rs)

	os.Remove(socket)
	l, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}

	go http.Serve(l, serverMux)

	return nil
}

func (a *Agent) GetSocketName() string {
	return a.socket
}

func (a *Agent) RunJob(ctx context.Context, program, driverPath string, appID int) (err error) {
	agentSock := a.GetSocketName()
	executorSock := fmt.Sprintf("/tmp/executorsock-%d", appID)

	cmd := exec.CommandContext(ctx, program,
		"--mode", "executor",
		"--agent-sock", agentSock,
		"--executor-sock", executorSock,
		"--driver-path", driverPath,
		"--app-id", strconv.Itoa(appID))

	// get the stderr log
	stderr, err := cmd.StderrPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stderr: %v", err)
		return
	}
	go func() {
		slurp, _ := ioutil.ReadAll(stderr)
		fmt.Printf("%s\n", string(slurp))
	}()

	// get the stdout log
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stdout: %v", err)
		return
	}
	go func() {
		slurp, _ := ioutil.ReadAll(stdout)
		fmt.Printf("%s\n", string(slurp))
	}()

	// start the cmd, does not wait for it to complete
	if err = cmd.Start(); err != nil {
		err = fmt.Errorf("executor start: %v", err)
		return
	}

	// waiting for the executor rpc to be ready
	b := time.Now()
	client, err := waitForReady(ctx, executorSock, 5*time.Second)
	if err != nil {
		err = fmt.Errorf("rpc dial: %v", err)
		return
	}
	a.logger.Infow("local executor is ready", "startup", time.Now().Sub(b))

	a.logger.Infow("local executor to run driver")

	req := true
	res := new(bool)
	if err = client.Call("Executor.Start", &req, &res); err != nil {
		err = fmt.Errorf("rpc start: %v", err)
		return
	}

	a.logger.Infow("local executor is shutting down")
	terReq := 0
	terRes := new(bool)
	// ignore error, since when the executor is terminated, this rpc will fail
	_ = client.Call("Executor.Terminate", &terReq, &terRes)

	if err = cmd.Wait(); err != nil {
		a.logger.Errorw("executor wait", "err", err)
		return
	}

	return
}

func waitForReady(ctx context.Context, executorSock string, expiredIn time.Duration) (
	*rpc.Client, error,
) {
	timeout := time.After(expiredIn)
	sleep := 10 * time.Millisecond
	for {
		time.Sleep(sleep)

		select {
		case <-ctx.Done():
			return nil, errors.New("cancel")
		case <-timeout:
			return nil, errors.New("timeout")
		default:
			client, err := rpc.DialHTTP("unix", executorSock)
			if err != nil {
				continue
			}
			return client, nil
		}
	}
}
