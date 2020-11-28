package md

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const metadataKeyUserUUID string = "x-user-uuid"

func AddUserUUIDToContext(ctx context.Context, uid string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyUserUUID, uid)
}
