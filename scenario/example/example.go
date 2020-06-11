package main

import (
	"log"
	"sync"
	"time"

	"github.com/gobench-io/gobench/scenario"
)

func Export() scenario.Vus {
	return scenario.Vus{
		scenario.Vu{
			Nu:   10,
			Rate: 100,
			Fu:   F,
		},
	}
}

func F(i int, donewg *sync.WaitGroup) {
	defer donewg.Done()

	count := 0

	for {
		count++
		if count > 10 {
			break
		}
		log.Printf("sub num %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
