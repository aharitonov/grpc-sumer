package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example.net/developer/grpc-sumer/sumerapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultVal = 123
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	x    = flag.Int("x", defaultVal, "X value")
	y    = flag.Int("y", defaultVal, "Y value")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSumerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.SumerRequest{X: int32(*x), Y: int32(*y)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResult())
}
