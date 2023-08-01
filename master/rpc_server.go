package master

import (
	"context"

	api "github.com/gobench-io/gobench/v2/gen/go/pb"
)

var _ api.JobDistributionServiceServer = (*Master)(nil)

func (m *Master) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	//TODO implement me

	return &api.PingResponse{}, nil
}
