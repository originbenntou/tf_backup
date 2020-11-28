package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/trend/domain/model"
)

type HistoryRepository interface {
	FindHistoryByUserUuid(context.Context, string, uint64) (*model.History, error)
	CreateHistory(context.Context, *model.History) error
	FindHistoryWithSuggestInfoByUserUuid(context.Context, string) ([]*model.History, error)
}
