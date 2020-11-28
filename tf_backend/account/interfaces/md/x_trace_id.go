package md

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const metadataKeyTraceID string = "x-trace-id"

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyTraceID, traceID)
}

func GetTraceIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	values := md.Get(metadataKeyTraceID)
	if len(values) < 1 {
		return ""
	}
	return values[0]
}
