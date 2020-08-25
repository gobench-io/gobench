// a valid gobench scenario
// do nothing

package main

import (
	"context"

	"github.com/gobench-io/gobench/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   20,
			Rate: 100,
			Fu:   f1,
		},
	}
}

// this func does nothing
func f1(ctx context.Context, vui int) {}
