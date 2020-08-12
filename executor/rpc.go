package executor

import (
	"context"
	"os"
)

func (e *Executor) Start(req *bool, res *bool) (err error) {
	ctx := context.TODO()

	e.logger.Infow("executor rpc starting")

	err = e.driver.Run(ctx)

	e.logger.Infow("executor rpc finished")

	return
}

// Terminate shutdown this executor process
func (e *Executor) Terminate(req *int, res *bool) (err error) {
	os.Exit(*req)
	*res = true
	return
}
