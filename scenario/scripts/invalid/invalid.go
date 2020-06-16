// this scenario is invalid because there is no Export function
package main

import (
	"context"
	"log"
	"time"

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

// this function receive the ctx.Done signal
func f1(ctx context.Context, vui int) {
	for {
		log.Println("tic")
		time.Sleep(1 * time.Second)
	}
}
