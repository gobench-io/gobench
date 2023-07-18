//
// fan in scenario:
// A single subscriber reading from "prefix/clients/#" topic filter
// 1k publisher publishing to exclusive topic "prefix/clients/{client_id}"
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
			Nu:   1,
			Rate: 100,
			Fu:   subf,
		},
		{
			Nu:   1000,
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

	_ = client.Subscribe(ctx, "prefix/clients/#", 0, nil)
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
			_ = client.PublishToSelf(ctx, "prefix/clients/", 0, dis.RandomByte(150))
			dis.SleepRatePoisson(rate)
		}
	}
}
