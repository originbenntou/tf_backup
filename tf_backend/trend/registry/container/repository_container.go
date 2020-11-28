package container

import (
	"github.com/TrendFindProject/tf_backend/trend/domain/repository"
	repo "github.com/TrendFindProject/tf_backend/trend/infrastructure/datastore"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/mysql"
)

func (c Container) GetSearchRepository(db mysql.DBManager) repository.SearchRepository {
	return repo.NewSearchRepository(db)
}

func (c Container) GetHistoryRepository(db mysql.DBManager) repository.HistoryRepository {
	return repo.NewHistoryRepository(db)
}

func (c Container) GetSuggestRepository(db mysql.DBManager) repository.SuggestRepository {
	return repo.NewSuggestRepository(db)
}
