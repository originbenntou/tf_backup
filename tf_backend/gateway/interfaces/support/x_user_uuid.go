package support

import (
	"context"
)

type contextKeyUser struct{}

func AddUserToContext(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, uid)
}

func GetUserFromContext(ctx context.Context) string {
	uid, ok := ctx.Value(contextKeyUser{}).(string)
	if !ok {
		return ""
	}
	return uid
}
