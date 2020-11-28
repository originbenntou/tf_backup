package md

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const metadataKeyTraceID string = "x-trace-id"

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyTraceID, traceID)
}
