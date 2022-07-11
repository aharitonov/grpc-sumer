package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.net/developer/grpc-sumer/sumerapi"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSumerServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumerRequest) (*pb.SumerResponse, error) {
	log.Printf("Received: %v, %v", in.GetX(), in.GetY())
	summa := in.GetX() + in.GetY()
	return &pb.SumerResponse{Result: fmt.Sprintf("Summa: %v", summa)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSumerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
