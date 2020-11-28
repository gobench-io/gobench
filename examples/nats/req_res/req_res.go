package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gobench-io/gobench/clients/nats"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/executor/scenario"
)

var server = "127.0.0.1"

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   2,
			Rate: 100,
			Fu:   resf,
		},
		{
			Nu:   10,
			Rate: 100,
			Fu:   reqf,
		},
	}
}

func resf(ctx context.Context, vui int) {
	client, err := nats.NewNatClient(ctx, server)
	if err != nil {
		log.Println(err)
		return
	}
	err = client.QueueSubscribe(ctx, "rpc.*", "foo", func(msg *nats.Msg) {
		log.Printf("[vu %d]: %s\n", vui, string(msg.Data))
		msg.Respond([]byte(fmt.Sprintf("response from vui %d for request %s", vui, string(msg.Data))))
	})

	if err != nil {
		log.Println(err)
		return
	}
}

func reqf(ctx context.Context, vui int) {
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
			_ = client.Request(ctx, "rpc."+strconv.Itoa(vui),
				[]byte("hello world"+strconv.Itoa(i)), 2*time.Second)
			i++
			dis.SleepRateLinear(rate)
		}
	}
}
