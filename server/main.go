package main

import (
	"context"
	pb "go-grpc-gateway-openapi/api/gen/v1"
	"go-grpc-gateway-openapi/secret"
	"google.golang.org/grpc"
	"log"
	"net"
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

	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	pb.RegisterHelloWorldServer(s, &server{})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	err = secret.Middleware(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
