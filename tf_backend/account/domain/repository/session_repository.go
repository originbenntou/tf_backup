package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
)

type SessionRepository interface {
	FindExistTokenByUserUuid(context.Context, string) (string, error)
	CreateSession(context.Context, *model.Session) error
	DeleteSessionByUserUuid(context.Context, string) error
	CountValidSessionByCompanyId(context.Context, uint64) (uint64, error)
}
