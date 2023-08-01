// Test a server running on a local machine on port 8080.
// Send 10 requests per second for 2 minute from 5 nodes in parallel,
// which totals up to 50 requests per second altogether.

package main

import (
	"context"
	"log"
	"time"

	httpClient "github.com/gobench-io/gobench/v2/clients/http"
	"github.com/gobench-io/gobench/v2/dis"
	"github.com/gobench-io/gobench/v2/executor/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   5,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	client, err := httpClient.NewHttpClient(ctx, "home")
	if err != nil {
		log.Println("create new client fail: " + err.Error())
		return
	}

	url := "http://localhost:8080/healthz"

	timeout := time.After(2 * time.Minute)

	for {
		select {
		case <-timeout:
			return
		default:
			go client.Get(ctx, url, nil)
			dis.SleepRatePoisson(10)
		}
	}
}
