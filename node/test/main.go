package main

import (
	"log"

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

	n.Run()
}
