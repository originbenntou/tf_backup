package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
)

type UserRepository interface {
	FindUserByEmail(context.Context, string) (*model.User, error)
	FindUserByUuid(context.Context, string) (*model.User, error)
	CreateUser(context.Context, *model.User) (uint64, error)
	CountUsersByCompanyId(context.Context, uint64) (uint64, error)
	CountUsersByUuid(context.Context, string) (uint64, error)
	ModifyUserPassword(context.Context, string, []byte) error
}
