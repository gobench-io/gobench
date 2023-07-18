package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gobench-io/gobench/v2/clients/nats"
	"github.com/gobench-io/gobench/v2/dis"
	"github.com/gobench-io/gobench/v2/executor/scenario"
)

var server = "127.0.0.1"

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   10,
			Rate: 100,
			Fu:   subf,
		},
		{
			Nu:   1,
			Rate: 100,
			Fu:   pubf,
		},
	}
}

func subf(ctx context.Context, vui int) {
	client, err := nats.NewNatClient(ctx, server)
	if err != nil {
		log.Println(err)
		return
	}
	err = client.QueueSubscribe(ctx, "hello.*", "foo", func(msg *nats.Msg) {
		log.Printf("[vu %d]: %s\n", vui, string(msg.Data))
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func pubf(ctx context.Context, vui int) {
	client, err := nats.NewNatClient(ctx, server)
	if err != nil {
		log.Println(err)
		return
	}

	rate := 1.0 // rps
	timeout := time.After(1 * time.Minute)

	i := 1

	for {
		select {
		case <-ctx.Done():
			break
		case <-timeout:
			_ = client.Disconnect(ctx)
			break
		default:
			_ = client.Publish(ctx, "hello."+strconv.Itoa(vui), []byte("hello world "+strconv.Itoa(i)))
			i++
			dis.SleepRateLinear(rate)
		}
	}
}
