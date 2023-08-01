package executor

import (
	"context"
	"os"

	api "github.com/gobench-io/gobench/v2/gen/go/pb"
)

// Start begins to run the program
func (e *Executor) Start(ctx context.Context, req *api.StartRequest) (*api.StartResult, error) {
	e.logger.Infow("executor rpc starting")

	err := e.run(ctx)

	e.logger.Infow("executor rpc finished")

	if err != nil {
		return nil, err
	}

	res := new(api.StartResult)
	res.AppID = int64(e.appID)
	res.Success = true

	return res, nil
}

// Terminate shutdown this executor process
func (e *Executor) Terminate(ctx context.Context, req *api.TermRequest) (*api.TermResult, error) {
	os.Exit(int(req.Code))

	res := new(api.TermResult)
	res.AppID = int64(e.appID)
	res.Success = true

	return res, nil
}
