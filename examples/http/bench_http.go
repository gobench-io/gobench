package main

// This runs a benchmark for 30 seconds, using 12 threads

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gobench-io/gobench"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
	httpClient "github.com/gobench-io/gobench/workers/http"

	"github.com/gobench-io/gobench/services/smtp"
)

func main() {
	bench := gobench.NewBench()

	// when the benchmark finishes, the result will be sent to these emails
	bench.
		Email(smtp.Config{
			Enable:   false,
			Host:     "smtp.sendgrid.net",
			Port:     25,
			Username: "apikey",
			Password: "insert-your-key",
			From:     "gobench <no-reply@veriksystems.com>",
			To:       []string{"dinh.nguyen@veriksystems.com"},
		}).
		Name("http benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	// must call after bench.Start
	go web.Serve(bench, 3001)

	// collect internal metrics
	go benchclient.InternalMonitor()

	// begin the bench
	vu := 12

	var donewg sync.WaitGroup
	donewg.Add(vu)

	var rate float64 = 1000 // per second
	for i := 0; i < vu; i++ {
		gobench.SleepPoisson(rate)

		go vuPool(i, &donewg)
	}

	donewg.Wait()

	time.Sleep(1 * time.Second)

	if err := bench.Finish(); err != nil {
		log.Printf("finish error %v\n", err)
	}
}

func vuPool(i int, donewg *sync.WaitGroup) {
	defer donewg.Done()

	ctx := context.Background()

	client1, err := httpClient.NewHttpClient(&ctx, "home")
	if err != nil {
		log.Println("create new client1 fail: " + err.Error())
		return
	}

	url1 := "http://192.168.2.29"

	headers := map[string]string{
		// "Content-Type": "application/json",
	}
	go func() {
		for {
			client1.Get(&ctx, url1, headers)
		}
	}()

	gobench.SleepLinear(1 / 30.0) // sleep for 30 seconds
}
