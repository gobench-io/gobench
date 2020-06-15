package main

import (
	"log"
	"time"

	"github.com/gobench-io/gobench/node"
)

func main() {
	so := "./scenario/example/example.so"

	n, err := node.New()
	if err != nil {
		log.Fatalln(err)
	}

	if err := n.Load(so); err != nil {
		log.Fatalln(err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		log.Printf("is running: %v\n", n.Running())
		time.Sleep(6 * time.Second)
		n.Cancel()
		log.Printf("is running: %v\n", n.Running())
	}()

	n.Run()
	time.Sleep(10 * time.Second)
}
