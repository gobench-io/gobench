package executor

import (
	"context"
	"os"

	"github.com/gobench-io/gobench/v2/pb"
)

// Start begins to run the program
func (m *Executor) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResult, error) {
	m.logger.Infow("executor rpc starting")

	err := m.run(ctx)

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
