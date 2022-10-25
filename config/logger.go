package config

import (
	"os"

	"google.golang.org/grpc/grpclog"
)

var GrpcLog grpclog.LoggerV2

func SetupLogger() {
	GrpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(GrpcLog)
}
