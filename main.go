package main

import (
	"log"

	"github.com/gobench-io/gobench/server"
	"github.com/gobench-io/gobench/web"
	"github.com/gobench-io/gobench/workers/benchclient"
)

func main() {
	bench := server.NewBench()

	bench.Name("mqtt fan out benchmark example")

	if err := bench.Start(); err != nil {
		log.Fatalln(err)
	}

	go benchclient.InternalMonitor()
	web.Serve(bench, 3001)
}
