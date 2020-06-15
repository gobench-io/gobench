package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/scenario"

	httpClient "github.com/gobench-io/gobench/workers/http"
)

func Export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   2,
			Rate: 100,
			Fu:   f1,
		},
		scenario.Vu{
			Nu:   1,
			Rate: 100,
			Fu:   f2,
		},
	}
}

func f1(ctx context.Context, vui int) {
	count := 0

	for {
		count++
		if count > 20 {
			break
		}
		log.Printf("sub num %d\n", vui)
		time.Sleep(1 * time.Second)
	}
}

func f2(ctx context.Context, vui int) {
	client1, err := httpClient.NewHttpClient(&ctx, "home")
	if err != nil {
		log.Println("create new client1 fail: " + err.Error())
		return
	}

	url1 := "http://192.168.2.35"

	headers := map[string]string{
		// "Content-Type": "application/json",
	}

	count := 0

	for {
		count++
		if count > 10 {
			break
		}

		client1.Get(&ctx, url1, headers)
		time.Sleep(1 * time.Second)
	}
}
