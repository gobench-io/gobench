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
	"sync"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	"github.com/gobench-io/gobench/workers/mqtt"
)

func main() {
	bench := gobench.NewBench()
	bench.Name("mqtt fan out benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go web.Serve(bench, 3001)
	go benchclient.InternalMonitor()

	subVu := 1000
	pubVu := 1

	var donewg sync.WaitGroup
	donewg.Add(pubVu + subVu)

	rate := 100.0 // per second

	for i := 0; i < subVu; i++ {
		gobench.SleepPoisson(rate)

		go subVuPool(i, &donewg)
	}

	for j := 0; j < pubVu; j++ {
		gobench.SleepPoisson(rate)

		go pubVuPool(j, &donewg)
	}

	donewg.Wait()

	if err := bench.Finish(); err != nil {
		log.Printf("finish error %v\n", err)
	}
}

func subVuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()

	defer donewg.Done()

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

func pubVuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()

	defer donewg.Done()

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
		gobench.SleepPoisson(rate)
		go func() {
			_ = client.Publish(&ctx, "fixed/broadcast/topic", 0, gobench.RandomByte(150))
		}()
	}

	// finally
	_ = client.Disconnect(&ctx)
}
