//
// query scenario:
// This scenario simulates client server communication pattern.
// 1k servers subscribes for private topic and replies for incoming requests.
// 1 client sends request to one random server each time with 500 rps frequency.
// Request is sent with QOS2, response is delivered with QOS1.
// Overall Msg rate: 1k msg/s
// Message Size: 150 random bytes
// Runtime: 5 min
//

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/gobench-io/gobench/v2/clients/mqtt"
	"github.com/gobench-io/gobench/v2/dis"
	"github.com/gobench-io/gobench/v2/executor/scenario"
)

const (
	clientNum = 1
	serverNum = 1000
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   1,
			Rate: 100,
			Fu:   clientf,
		},
		{
			Nu:   1000,
			Rate: 100,
			Fu:   serverf,
		},
	}
}

func clientf(ctx context.Context, vui int) {
	clientID := fmt.Sprintf("client-%d", vui)

	opts := mqtt.NewClientOptions()
	opts.
		AddBroker("192.168.2.35:1883").
		SetClientID(clientID)

	client, err := mqtt.NewMqttClient(ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(ctx); err != nil {
		log.Println(err)
		return
	}

	if err = client.Subscribe(ctx, "prefix/clients/#", 1, nil); err != nil {
		log.Println(err)
		return
	}

	rate := 500.0 // rps
	timeout := time.After(5 * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return
		case <-timeout:
			_ = client.Disconnect(ctx)
			return
		default:
			topic := fmt.Sprintf("prefix/servers/server-%d", rand.Intn(serverNum))
			_ = client.Publish(ctx, topic, 2, dis.RandomByte(150))
			dis.SleepRatePoisson(rate)
		}
	}
}

func serverf(ctx context.Context, vui int) {
	clientID := fmt.Sprintf("server-%d", vui)

	opts := mqtt.NewClientOptions()
	opts.
		AddBroker("192.168.2.35:1883").
		SetClientID(clientID)

	client, err := mqtt.NewMqttClient(ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(ctx); err != nil {
		log.Println(err)
		return
	}

	if err = client.SubscribeToSelf(
		ctx,
		"prefix/servers/",
		2,
		func(c paho.Client, m paho.Message) {
			_ = client.PublishToSelf(ctx, "prefix/clients/", 1, m.Payload())
		},
	); err != nil {
		log.Println(err)
		return
	}
}
