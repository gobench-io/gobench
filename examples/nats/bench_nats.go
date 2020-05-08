package main

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	nats "github.com/gobench-io/gobench/workers/nats"
)

func main() {
	bench := gobench.NewBench()
	bench.Name("nats benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go web.Serve(bench, 3001)
	go benchclient.InternalMonitor()

	vu := 3000

	var donewg sync.WaitGroup
	donewg.Add(vu)

	// this is the pool
	var rate float64 = 1000 // per second
	for i := 0; i < vu; i++ {
		gobench.SleepPoisson(rate)
		go vuPool(i, &donewg)
	}

	donewg.Wait()

	gobench.SleepLinear(1.0)

	if err := bench.Finish(); err != nil {
		log.Printf("finish error %s\n", err.Error())
	}
}

func vuPool(i int, donewg *sync.WaitGroup) {
	ctx := context.Background()
	// nats opts
	// url := nats.DefaultURL
	defer donewg.Done()

	url := "127.0.0.1"
	client, err := nats.NewNatClient(&ctx, url)
	if err != nil {
		log.Println(err)
		return
	}

	// wait 1 sec
	gobench.SleepLinear(1.0)

	// subscribe_to_self("prefix/clients/", 0)
	_ = client.Subscribe(&ctx, "hello."+strconv.Itoa(i))

	// loop(time = 5 min, rate = 1 rps)
	rate := 1.0
	for j := 0; j < 10; j++ {
		gobench.SleepPoisson(rate)
		_ = client.Publish(&ctx, "hello."+strconv.Itoa(i), []byte("hello world"))
	}
	// finally
	_ = client.Disconnect(&ctx)
}
