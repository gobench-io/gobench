//
// fan out (broadcast) scenario:
// 1k subscribers reading from the same topic "fixed/broadcast/topic"
// 1 publisher sending 1 msg/s to topic "fixed/broadcast/topic"
// Overall Msg rate: 1k msg/s
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
			Nu:   1000,
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

	_ = client.Subscribe(ctx, "fixed/broadcast/topic", 0, nil)
}

func pubf(ctx context.Context, vui int) {
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
			_ = client.Publish(ctx, "fixed/broadcast/topic", 0, dis.RandomByte(150))
			dis.SleepRatePoisson(rate)
		}
	}
}
