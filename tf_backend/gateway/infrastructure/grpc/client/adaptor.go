package client

import (
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func GetGrpcConn(target string, interceptors ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
	chain := grpc_middleware.ChainUnaryClient(interceptors...)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithUnaryInterceptor(chain))
	if err != nil {
		logger.Common.Error(err.Error())
	}
	return conn
}
