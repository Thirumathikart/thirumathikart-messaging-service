package middlewares

import (
	"context"
	"time"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	var h interface{}
	var err error
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		h, err = nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	} else {
		clientSecrets := md["secret"]
		serverSecret := config.ServerSecret
		if clientSecrets[0] != serverSecret {
			h, err = nil, status.Errorf(codes.InvalidArgument, "missing 'x-request-id' header")
		} else {
			h, err = handler(ctx, req)
		}
	}
	// Logging with grpclog (grpclog.LoggerV2)
	config.GrpcLog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
