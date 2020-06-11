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

	var donewg sync.WaitGroup

	var totalVu int

	for i := range vus {
		totalVu += vus[i].Nu
	}

	donewg.Add(totalVu)

	for i := range vus {
		for j := 0; j < vus[i].Nu; j++ {
			go func(j int) {
				vus[i].Fu(j, &donewg)
			}(j)
		}
	}

	donewg.Wait()
}
