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

	vus[0].Fu(1, &donewg)
}
