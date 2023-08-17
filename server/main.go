package main

import (
	"context"
	"log"
	"net"

	pb "go-grpc-gateway-openapi/api/gen/v1"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type server struct {
	pb.UnimplementedHelloWorldServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := "Hello, " + req.Name
	return &pb.HelloResponse{Message: message}, nil
}

func (s *server) PostMessage(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {
	result := "Posted: " + req.Content
	return &pb.PostResponse{Result: result}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
