package secret

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Middleware(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Metadata not provided")
	}

	tokenStrings := md.Get("authorization")
	if len(tokenStrings) == 0 {
		return status.Errorf(codes.Unauthenticated, "Authorization token not provided")
	}

	tokenString := tokenStrings[0]
	if tokenString == "" {
		return status.Errorf(codes.Unauthenticated, "Authorization token not provided")
	}

	if token := "ANKIT_KUMAR"; token != tokenString {
		return status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	return nil
}
