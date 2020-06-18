package main

// This runs a benchmark for 30 seconds, using 12 threads

import (
	"context"
	"log"
	"time"

	httpClient "github.com/gobench-io/gobench/clients/http"
	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   12,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	client1, err := httpClient.NewHttpClient(&ctx, "home")
	if err != nil {
		log.Println("create new client1 fail: " + err.Error())
		return
	}

	url1 := "http://192.168.2.29"

	headers := map[string]string{
		// "Content-Type": "application/json",
	}
	go func() {
		for {
			client1.Get(&ctx, url1, headers)
		}
	}()

	time.Sleep(30 * time.Second)
}
