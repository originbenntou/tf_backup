package interceptor

import (
	"context"
	"fmt"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/md"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func XTraceID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		traceID := md.GetTraceIDFromContext(ctx)
		ctx = md.AddTraceIDToContext(ctx, traceID)
		return handler(ctx, req)
	}
}

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		h, err := handler(ctx, req)
		ctxzap.AddFields(ctx,
			zap.String("TraceID", md.GetTraceIDFromContext(ctx)),
			zap.String("UserUUID", fmt.Sprintf("%s", md.GetUserUUIDFromContext(ctx))),
		)
		return h, err
	}
}

func XUserUUID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		uid, err := md.SafeGetUserUUIDFromContext(ctx)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		ctx = md.AddUserUUIDToContext(ctx, uid)
		return handler(ctx, req)
	}
}
