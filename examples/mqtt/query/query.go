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
	"sync"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	"github.com/gobench-io/gobench/workers/mqtt"
)

func main() {
	bench := gobench.NewBench()
	bench.Name("mqtt fan in benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go web.Serve(bench, 3001)
	go benchclient.InternalMonitor()

	clientVu := 1
	serverVu := 1000

	var donewg sync.WaitGroup
	donewg.Add(serverVu + clientVu)

	rate := 1000.0 // per second

	for i := 0; i < clientVu; i++ {
		gobench.SleepPoisson(rate)

		go clientVuPool(i, &donewg)
	}

	for j := 0; j < serverVu; j++ {
		gobench.SleepPoisson(rate)

		go serverVuPool(j, &donewg)
	}

	donewg.Wait()

	if err := bench.Finish(); err != nil {
		log.Printf("finish error %v\n", err)
	}
}

func clientVuPool(i int, donewg *sync.WaitGroup) {
	defer donewg.Done()

	ctx := context.Background()

	clientID := fmt.Sprintf("client-%d", i)

	opts := mqtt.NewClientOptions()
	opts.
		AddBroker("192.168.2.29:1883").
		SetClientID(clientID)

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	if err = client.Subscribe(&ctx, "prefix/clients/#", 1, nil); err != nil {
		log.Println(err)
		return
	}

	rate := 500.0 // rps
	for j := 0; j < int(60*5*rate); j++ {
		gobench.SleepPoisson(rate)

		topic := fmt.Sprintf("prefix/servers/server-%d", rand.Intn(1000))
		_ = client.Publish(&ctx, topic, 2, gobench.RandomByte(150))
	}

	// finally
	_ = client.Disconnect(&ctx)
}

func serverVuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()

	defer donewg.Done()

	clientID := fmt.Sprintf("server-%d", i)

	opts := mqtt.NewClientOptions()
	opts.
		AddBroker("192.168.2.29:1883").
		SetClientID(clientID)

	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	if err = client.SubscribeToSelf(
		&ctx,
		"prefix/servers/",
		2,
		func(c paho.Client, m paho.Message) {
			_ = client.PublishToSelf(&ctx, "prefix/clients/", 1, m.Payload())
		},
	); err != nil {
		log.Println(err)
		return
	}
}
