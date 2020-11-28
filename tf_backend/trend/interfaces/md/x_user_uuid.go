package md

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

const metadataKeyUserUUID string = "x-user-uuid"

func AddUserUUIDToContext(ctx context.Context, uid string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKeyUserUUID, uid)
}

var ErrNotFoundUserUUID = errors.New("not found user uuid")

func GetUserUUIDFromContext(ctx context.Context) string {
	uid, err := SafeGetUserUUIDFromContext(ctx)
	if err != nil {
		panic(err)
	}
	return uid
}

func SafeGetUserUUIDFromContext(ctx context.Context) (uid string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uid, ErrNotFoundUserUUID
	}
	values := md.Get(metadataKeyUserUUID)
	if len(values) < 1 {
		return uid, ErrNotFoundUserUUID
	}
	return values[0], nil
}
