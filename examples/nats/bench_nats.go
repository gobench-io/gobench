package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/scenario"
	nats "github.com/gobench-io/gobench/workers/nats"
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
	client, err := nats.NewNatClient(&ctx, url)
	if err != nil {
		log.Println(err)
		return
	}

	// wait 1 sec
	dis.SleepRatePoisson(1)

	// subscribe_to_self("prefix/clients/", 0)
	_ = client.Subscribe(&ctx, "hello."+strconv.Itoa(vui))

	// loop(time = 5 min, rate = 1 rps)
	rate := 1.0
	for j := 0; j < 10; j++ {
		dis.SleepRatePoisson(rate)
		_ = client.Publish(&ctx, "hello."+strconv.Itoa(vui), []byte("hello world"))
	}
	// finally
	_ = client.Disconnect(&ctx)
}
