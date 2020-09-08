package main

// This runs a benchmark for 30 seconds, using 12 threads

import (
	"context"
	"log"
	"time"

	httpClient "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/executor/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   12,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	client1, err := httpClient.NewHttpClient(ctx, "home")
	if err != nil {
		log.Println("create new client1 fail: " + err.Error())
		return
	}

	url1 := "http://192.168.2.35"

	headers := map[string]string{
		// "Content-Type": "application/json",
	}

	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeout:
			return
		default:
			client1.Get(ctx, url1, headers)
		}
	}
}
