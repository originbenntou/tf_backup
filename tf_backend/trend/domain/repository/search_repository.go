package repository

import (
	"context"
	"time"

	"github.com/TrendFindProject/tf_backend/trend/domain/model"
)

type SearchRepository interface {
	FindSearchBySearchWord(context.Context, string, time.Time) (*model.Search, error)
	CreateSearch(context.Context, *model.Search) (uint64, error)
	FindSearchById(context.Context, uint64) (*model.Search, error)
}
