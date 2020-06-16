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

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/scenario"
	"github.com/gobench-io/gobench/workers/mqtt"
)

func Export() scenario.Vus {
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

	// wait for all other workers finish the connect step
	poolSignal.Done()
	poolSignal.Wait()

	_ = client.SubscribeToSelf(&ctx, "prefix/clients/", 0, nil)

	rate := 1.0 // rps
	for j := 0; j < 60*5; j++ {
		dis.SleepRatePoisson(rate)
		_ = client.PublishToSelf(&ctx, "prefix/clients/", 0, gobench.RandomByte(150))
	}

	// finally
	_ = client.Disconnect(&ctx)
}
