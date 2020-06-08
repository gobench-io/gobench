package main

import (
	"log"

	"github.com/gobench-io/gobench/server"
	"github.com/gobench-io/gobench/web"
)

// func main() {
// 	bench := server.NewBench()

// 	bench.Name("mqtt fan out benchmark example")

// 	if err := bench.Start(); err != nil {
// 		log.Fatalln(err)
// 	}

// 	go benchclient.InternalMonitor()
// 	web.Serve(bench, 3001)
// }

func main() {
	server, _ := server.New()

	if err := server.Start(); err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	web.Serve(server, 3001)
}
