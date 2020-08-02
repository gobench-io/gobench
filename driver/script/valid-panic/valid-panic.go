// a valid gobench scenario
// the vu function panic

package main

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

// this function panic
func f1(ctx context.Context, vui int) {
	y := 0
	x := 1 / y
	_ = x
}
