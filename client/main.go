package main

import (
	"context"
	"log"

	pb "go-grpc-gateway-openapi/api/gen/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000" // Update this with the server address
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	// SayHello RPC call
	name := "Ankit"
	helloReq := &pb.HelloRequest{Name: name}
	helloResp, err := c.SayHello(context.Background(), helloReq)
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
