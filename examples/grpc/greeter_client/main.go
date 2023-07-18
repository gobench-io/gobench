package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/clients/gbGrpc"
	"github.com/gobench-io/gobench/v2/executor/scenario"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func export() scenario.Vus {
	return scenario.Vus{
		{
			Nu:   5,
			Rate: 1000,
			Fu:   f,
		},
	}
}

func f(ctx context.Context, vui int) {
	address := "localhost:50051"

	// Set up a connection to the server.
	conn, err := gbGrpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := fmt.Sprintf("hello from vui %d", vui)

	timeout := time.After(2 * time.Minute)

	for {
		select {
		case <-timeout:
			return

		default:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("Could not greet: %v\n", err)
			}
			log.Printf("Greeting: %s", r.GetMessage())
		}
	}
}
