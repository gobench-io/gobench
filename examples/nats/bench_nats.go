package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gobench-io/gobench/clients/nats"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/executor/scenario"
)

func Export() scenario.Vus {
	// nats benchmark example
	return scenario.Vus{
		{
			Nu:   3000,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	url := "127.0.0.1"
	client, err := nats.NewNatClient(ctx, url)
	if err != nil {
		log.Println(err)
		return
	}

	// wait 1 sec
	dis.SleepRatePoisson(1)

	// subscribe_to_self("prefix/clients/", 0)
	_ = client.Subscribe(ctx, "hello."+strconv.Itoa(vui))

	rate := 1.0 // rps
	timeout := time.After(5 * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-timeout:
			_ = client.Disconnect(ctx)
			return
		default:
			_ = client.Publish(ctx, "hello."+strconv.Itoa(vui), []byte("hello world"))
			dis.SleepRatePoisson(rate)
		}
	}
}
