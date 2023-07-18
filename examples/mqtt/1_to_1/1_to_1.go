//
// 1_to_1 Scenario:
// 2k clients subscribe to an exclusive topic: "prefix/clients/{clientID}"
// The same 2k clients send messages on that topic to themselves
// Overall Msg rate: 2k msg/s
// Message Size: 150 random bytes
// Runtime: 5 min
//

package main

import (
	"context"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/clients/mqtt"
	"github.com/gobench-io/gobench/v2/dis"
	"github.com/gobench-io/gobench/v2/executor/scenario"
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   2000,
			Rate: 100,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("192.168.2.35:1883")

	client, err := mqtt.NewMqttClient(ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(ctx); err != nil {
		log.Println(err)
		return
	}

	_ = client.SubscribeToSelf(ctx, "prefix/clients/", 0, nil)

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
			_ = client.PublishToSelf(ctx, "prefix/clients/", 0, dis.RandomByte(150))
			dis.SleepRatePoisson(rate)
		}
	}
}
