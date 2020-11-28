package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/trend/domain/model"
)

type SuggestRepository interface {
	FindSuggestsBySearchId(context.Context, uint64) ([]*model.Suggest, error)
}
