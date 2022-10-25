package middlewares

import (
	"context"
	"time"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	"google.golang.org/grpc"
)

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Calls the handler
	h, err := handler(ctx, req)

	// Logging with grpclog (grpclog.LoggerV2)
	config.GrpcLog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
