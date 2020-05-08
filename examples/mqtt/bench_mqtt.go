package main

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	"github.com/gobench-io/gobench/workers/mqtt"
)

func main() {
	bench := gobench.NewBench()
	bench.Name("mqtts benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go web.Serve(bench, 3001)
	go benchclient.InternalMonitor()

	vu := 1000

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
		log.Printf("finish error %s\n", err.Error())
	}
}

func vuPool(i int, donewg, poolSignal *sync.WaitGroup) {
	ctx := context.Background()

	// mqtt opts
	defer donewg.Done()

	host := "127.0.0.1:1883"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(host)
	client, err := mqtt.NewMqttClient(&ctx, opts)
	if err != nil {
		log.Println(err)
		return
	}
	// wait for all worker finish the connect step
	poolSignal.Done()
	poolSignal.Wait()

	if err = client.Connect(&ctx); err != nil {
		log.Println(err)
		return
	}

	// wait 1 sec
	gobench.SleepLinear(1)

	// subscribe_to_self("prefix/clients/", 0)
	_ = client.Subscribe(&ctx, "hello/"+strconv.Itoa(i), 0)

	// loop(time = 5 min, rate = 1 rps)
	rate := 1.0
	for j := 0; j < 60; j++ {
		gobench.SleepLinear(rate)
		_ = client.Publish(&ctx, "hello/"+strconv.Itoa(i), 0, []byte("hello world"))
	}
	// finally
	_ = client.Disconnect(&ctx)
}
