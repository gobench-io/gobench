package agent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/pb"
	"google.golang.org/grpc"
)

type Options struct {
	Route       string
	ClusterPort int
	Socket      string
}

// Agent struct
// todo: agent needs appID
type Agent struct {
	mu sync.Mutex

	route       string
	clusterPort int

	ml             pb.AgentServer
	logger         logger.Logger
	executorLogger io.WriteCloser // when running the executor, save log here
	socket         string         // unix socket that the agent rpc server will listen at
}

func NewLocalAgent(ml pb.AgentServer, logger logger.Logger) (*Agent, error) {
	a := &Agent{
		ml:     ml,
		logger: logger,
	}
	return a, nil
}

func NewAgent(opts *Options, ml pb.AgentServer, logger logger.Logger) (*Agent, error) {
	a := &Agent{
		route:       opts.Route,
		clusterPort: opts.ClusterPort,
		socket:      opts.Socket,
		logger:      logger,
		ml:          ml,
	}
	return a, nil
}

// SetMetricLogger sets metric logger property
func (a *Agent) SetMetricLogger(ml pb.AgentServer) {
	a.mu.Lock()
	a.ml = ml
	a.mu.Unlock()
}

// SetExecutorLogger sets executor log writer property
func (a *Agent) SetExecutorLogger(w io.WriteCloser) *Agent {
	a.mu.Lock()
	a.executorLogger = w
	a.mu.Unlock()

	return a
}

// SetLogger set the logger property
func (a *Agent) SetLogger(l logger.Logger) {
	a.mu.Lock()
	a.logger = l
	a.mu.Unlock()
}

// StartSocketServer setup an rpc server over agent unix socket
// the function runs the server in a separate routine
func (a *Agent) StartSocketServer() error {
	// remove if any
	os.Remove(a.socket)

	l, err := net.Listen("unix", a.socket)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterAgentServer(s, a.ml)

	go s.Serve(l)

	return nil
}

// RunJob runs the executor in a shell
func (a *Agent) RunJob(ctx context.Context, executorPath string, appID int) (err error) {
	agentSock := a.socket
	executorSock := fmt.Sprintf("/tmp/gobench-executorsock-%d-%d", appID, time.Now().Unix())

	cmd := exec.CommandContext(ctx, executorPath,
		"--agent-sock", agentSock,
		"--executor-sock", executorSock)

	// get the stderr log
	stderr, err := cmd.StderrPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stderr: %v", err)
		return
	}

	go func() {
		if _, err := io.Copy(a.executorLogger, stderr); err != nil {
			a.logger.Errorw("failed write executor log", "err", err)
		}
	}()

	// get the stdout log
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stdout: %v", err)
		return
	}
	go func() {
		if _, err := io.Copy(a.executorLogger, stdout); err != nil {
			a.logger.Errorw("failed write executor log", "err", err)
		}
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

	// todo: handle the response
	if _, err = client.Start(ctx, &pb.StartRequest{}); err != nil {
		err = fmt.Errorf("rpc start: %v", err)
		return
	}

	a.logger.Infow("local executor is shutting down")

	// ignore error, since when the executor is terminated, this rpc will fail
	_, _ = client.Terminate(ctx, &pb.TermRequest{})

	if err = cmd.Wait(); err != nil {
		a.logger.Errorw("executor wait", "err", err)
		return
	}

	return
}

func waitForReady(ctx context.Context, executorSock string, expiredIn time.Duration) (
	pb.ExecutorClient, error,
) {
	timeout := time.After(expiredIn)
	sleep := 10 * time.Millisecond
	socket := "passthrough:///unix://" + executorSock
	for {
		time.Sleep(sleep)

		select {
		case <-ctx.Done():
			return nil, errors.New("cancel")
		case <-timeout:
			return nil, errors.New("timeout")
		default:
			conn, err := grpc.Dial(socket, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				continue
			}
			client := pb.NewExecutorClient(conn)
			return client, nil
		}
	}
}
