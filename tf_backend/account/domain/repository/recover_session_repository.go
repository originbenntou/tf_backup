package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
)

type RecoverSessionRepository interface {
	CreateRecoverSession(context.Context, *model.RecoverSession) error
	FindRecoverSessionByUuid(context.Context, string) (*model.RecoverSession, error)
	DeleteRecoverSessionByUuid(context.Context, string) error
}
