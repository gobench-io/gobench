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

	"github.com/gobench-io/gobench/clients/mqtt"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
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
	opts.AddBroker("192.168.2.29:1883")

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

	// finally
	// _ = client.Disconnect(ctx)
}

func pubf(ctx context.Context, vui int) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("192.168.2.29:1883")

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
	for j := 0; j < int(60*5*rate); j++ {
		dis.SleepRatePoisson(rate)
		go func() {
			_ = client.PublishToSelf(ctx, "prefix/clients/", 0, dis.RandomByte(150))
		}()
	}

	// finally
	_ = client.Disconnect(ctx)
}
