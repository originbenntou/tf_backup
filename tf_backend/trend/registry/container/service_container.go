package container

import (
	"github.com/TrendFindProject/tf_backend/trend/application/service"
	"github.com/TrendFindProject/tf_backend/trend/domain/repository"
)

func (c Container) GetTrendService(
	tr repository.SearchRepository,
	sr repository.SuggestRepository,
	hr repository.HistoryRepository,
) service.TrendService {
	return service.NewTrendService(tr, sr, hr)
}
