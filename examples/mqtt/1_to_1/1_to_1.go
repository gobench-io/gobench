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
	"sync"
	"time"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	"github.com/gobench-io/gobench/workers/mqtt"
)

func main() {
	bench := gobench.NewBench()
	bench.Name("mqtt 1-to-1 benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go web.Serve(bench, 3001)
	go benchclient.InternalMonitor()

	vu := 2000

	var donewg, poolSignal sync.WaitGroup
	donewg.Add(vu)
	poolSignal.Add(vu)

	var rate float64 = 100 // per second
	for i := 0; i < vu; i++ {
		gobench.SleepPoisson(rate)

		go vuPool(i, &donewg, &poolSignal)
	}

	donewg.Wait()

	time.Sleep(1 * time.Second)

	if err := bench.Finish(); err != nil {
		log.Printf("finish error %v\n", err)
	}
}

func vuPool(i int, donewg, poolSignal *sync.WaitGroup) {
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

	// wait for all other workers finish the connect step
	poolSignal.Done()
	poolSignal.Wait()

	_ = client.SubscribeToSelf(&ctx, "prefix/clients/", 0, nil)

	rate := 1.0 // rps
	for j := 0; j < 60*5; j++ {
		gobench.SleepPoisson(rate)
		_ = client.PublishToSelf(&ctx, "prefix/clients/", 0, gobench.RandomByte(150))
	}

	// finally
	_ = client.Disconnect(&ctx)
}
