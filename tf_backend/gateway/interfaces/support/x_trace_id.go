package support

import (
	"context"
)

type contextKeyTraceID struct{}

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, contextKeyTraceID{}, traceID)
}

// 有効なTraceIDがない場合は空文字を記録する
func GetTraceIDFromContext(ctx context.Context) string {
	id := ctx.Value(contextKeyTraceID{})
	if id == nil {
		return ""
	}

	traceID, ok := id.(string)
	if !ok {
		return ""
	}
	return traceID
}
