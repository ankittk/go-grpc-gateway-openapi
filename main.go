package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "go-grpc-gateway-openapi/api/gen/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gatewayMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gatewayMux, "localhost:9000", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Serve Swagger UI using http-swagger package
	fs := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	mux := http.NewServeMux()
	mux.Handle("/", gatewayMux)                                    // Handle gRPC-gateway requests
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs)) // Handle Swagger UI requests

	log.Println("Serving gRPC-gateway and Swagger UI on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
