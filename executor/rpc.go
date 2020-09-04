package executor

import (
	"context"
	"os"

	"github.com/gobench-io/gobench/pb"
)

// RPC is the struct that expose RPC interface
// avoid the rpc.Register: method "xxx" has 1 input parameters; needs exactly three
type RPC struct {
	e *Executor
}

func newExecutorRPC(e *Executor) (er *RPC, err error) {
	return &RPC{
		e: e,
	}, nil
}

func (er *RPC) Start(req *bool, res *bool) (err error) {
	ctx := context.TODO()

	er.e.logger.Infow("executor rpc starting")

	err = er.e.driver.Run(ctx)

	er.e.logger.Infow("executor rpc finished")

	return
}

// Terminate shutdown this executor process
func (er *RPC) Terminate(req *int, res *bool) (err error) {
	os.Exit(*req)
	*res = true
	return
}

// Start begins to run the program
func (m *Executor) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResult, error) {
	m.logger.Infow("executor rpc starting")

	err := m.driver.Run(ctx)

	m.logger.Infow("executor rpc finished")

	if err != nil {
		return nil, err
	}

	res := new(pb.StartResult)
	res.AppID = int64(m.appID)
	res.Success = true

	return res, nil
}

// Terminate shutdown this executor process
func (m *Executor) Terminate(ctx context.Context, req *pb.TermRequest) (*pb.TermResult, error) {
	os.Exit(int(req.Code))

	res := new(pb.TermResult)
	res.AppID = int64(m.appID)
	res.Success = true

	return res, nil
}
