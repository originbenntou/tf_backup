package interceptor

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/interfaces/md"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		)
		return h, err
	}
}
