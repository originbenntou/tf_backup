package interceptor

import (
	"context"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/md"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/support"
	"google.golang.org/grpc"
)

func XTraceID(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	// grpcに渡すときはmetadataに変換する（理由はよくわかっていない）
	traceID := support.GetTraceIDFromContext(ctx)
	ctx = md.AddTraceIDToContext(ctx, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func XUserUUID(ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {
	uid := support.GetUserFromContext(ctx)
	ctx = md.AddUserUUIDToContext(ctx, uid)
	return invoker(ctx, method, req, reply, cc, opts...)
}
