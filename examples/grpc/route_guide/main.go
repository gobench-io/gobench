package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/gobench-io/gobench/v2/clients/gbGrpc"
	"github.com/gobench-io/gobench/v2/executor/scenario"
	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/route_guide/routeguide"
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
		log.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	timeout := time.After(2 * time.Minute)

	for {
		select {
		case <-timeout:
			return

		default:
			listFeatures(client)
		}
	}
}

func listFeatures(client pb.RouteGuideClient) {
	rect := &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Printf("%v.ListFeatures(_) = _, %v", client, err)
		return
	}

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
	}
}
