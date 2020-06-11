package main

import (
	"log"
	"sync"
	"time"

	"github.com/gobench-io/gobench/scenario"
)

func Export() *scenario.Vus {
	return &scenario.Vus{
		scenario.Vu{
			Number:   1000,
			Function: SubVuPool,
		},
	}
}

func SubVuPool(i int, donewg *sync.WaitGroup) {
	defer donewg.Done()

	for {
		log.Println("sub ...")
		time.Sleep(1 * time.Second)
	}

}
