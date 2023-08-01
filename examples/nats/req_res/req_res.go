package main

import (
	"context"
	"fmt"
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
		msg.Respond([]byte(fmt.Sprintf("response from vui %d for request %s", vui, string(msg.Data))))
	})

	if err != nil {
		log.Println(err)
		return
	}
	err = client.Flush()
	if err != nil {
		fmt.Println(err)
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

L:
	for {
		select {
		case <-ctx.Done():
			break L
		case <-timeout:
			_ = client.Disconnect(ctx)
			break L
		default:
			reqMsg := fmt.Sprintf("req vui %d. no %d", vui, i)
			m, err := client.Request(ctx, "rpc."+strconv.Itoa(vui), []byte(reqMsg), 2*time.Second)
			if err == nil {
				log.Println(string(m.Data))
			} else {
				log.Printf("req fail for %s: %s", reqMsg, err)
			}
			i++
			dis.SleepRateLinear(rate)
		}
	}
}
