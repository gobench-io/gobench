package main

//go:generate go build -buildmode=plugin -o valid.so valid.go

import (
	"context"

	"github.com/gobench-io/gobench/scenario"
)

// Export is a required function for a scenario
func Export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   20,
			Rate: 100,
			Fu:   f1,
		},
	}
}

// this function waiting to receive ctx.Done
func f1(ctx context.Context, vui int) {
	select {
	case <-ctx.Done():
	}
}
