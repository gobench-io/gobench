package main

//go:generate go build -buildmode=plugin -o valid.so valid.go

import (
	"context"
	"log"
	"time"

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

// this function receive the ctx.Done signal
func f1(ctx context.Context, vui int) {
	for {
		log.Println("tic")
		time.Sleep(1 * time.Second)
	}
}
