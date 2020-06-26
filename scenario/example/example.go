package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/scenario"

	httpClient "github.com/gobench-io/gobench/clients/http"
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

// this function receive the ctx.Done signal
func f1(ctx context.Context, vui int) {
	ch := make(chan struct{})

	go func(c chan struct{}) {
		count := 0
		for range time.Tick(1 * time.Second) {
			count++
			if count > 4 {
				close(c)
				return
			}
			c <- struct{}{}
		}
	}(ch)

	for {
		select {
		case _, more := <-ch:
			if !more {
				log.Printf("f1 task is done")
				return
			}
			log.Printf("sub num %d\n", vui)
		case <-ctx.Done():
			log.Printf("f1 asked to exit")
			return
		}
	}
}

func f2(ctx context.Context, vui int) {
	client1, err := httpClient.NewHttpClient(ctx, "home")
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

		client1.Get(ctx, url1, headers)
		time.Sleep(1 * time.Second)
	}
}
