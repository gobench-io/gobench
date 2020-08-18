package executor

import (
	"context"
	"os"
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
