package mw

import (
	"context"

	"google.golang.org/grpc"
)

// RecoverUnaryServerInterceptor ...
func RecoverUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// TODO defer recover
		return handler(ctx, req)
	}
}
