package main

import (
	"log"
	"sync"

	"github.com/gobench-io/gobench/scenario"
)

func main() {
	so := "./scenario/example/example.so"

	vus, err := scenario.LoadPlugin(so)

	if err != nil {
		log.Println(err)
	}

	log.Printf("vus: %+v\n", vus)

	var donewg sync.WaitGroup

	var totalVu int

	for i := range vus {
		totalVu += vus[i].Nu
	}

	donewg.Add(totalVu)

	for i:= range vus{
		go vus[i].Fu(i, &donewg)
	}

	donewg.Wait()
}
