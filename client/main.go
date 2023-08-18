package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"

	pb "go-grpc-gateway-openapi/api/gen/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000" // Update this with the server address
	token   = "ADD_YOUR_TOKEN" // Update this with your token
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", token))
	// SayHello RPC call
	name := "Ankit"
	helloReq := &pb.HelloRequest{Name: name}
	helloResp, err := c.SayHello(ctx, helloReq)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("SayHello Response: %s", helloResp.Message)

	// PostMessage RPC call
	content := "This is a post request."
	postReq := &pb.PostRequest{Content: content}
	postResp, err := c.PostMessage(context.Background(), postReq)
	if err != nil {
		log.Fatalf("Failed to call PostMessage: %v", err)
	}
	log.Printf("PostMessage Response: %s", postResp.Result)
}
