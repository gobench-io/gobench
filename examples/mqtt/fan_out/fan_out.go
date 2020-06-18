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

	"github.com/gobench-io/gobench/clients/mqtt"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
	// mqtt fan out benchmark example
	return scenario.Vus{
		{
			Nu:   1,
			Rate: 1000,
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
	opts.AddBroker("192.168.2.29:1883")

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	_ = client.Subscribe(&ctx, "fixed/broadcast/topic", 0, nil)

	// finally
	// _ = client.Disconnect(&ctx)
}

func pubf(ctx context.Context, vui int) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("192.168.2.29:1883")

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	rate := 1.0 // rps
	for j := 0; j < int(60*5*rate); j++ {
		dis.SleepRatePoisson(rate)
		go func() {
			_ = client.Publish(&ctx, "fixed/broadcast/topic", 0, dis.RandomByte(150))
		}()
	}

	// finally
	_ = client.Disconnect(&ctx)
}
