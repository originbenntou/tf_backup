package container

import (
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	repo "github.com/TrendFindProject/tf_backend/account/infrastructure/datastore"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
)

func (c Container) GetUserRepository(db mysql.DBManager) repository.UserRepository {
	return repo.NewUserRepository(db)
}

func (c Container) GetCompanyRepository(db mysql.DBManager) repository.CompanyRepository {
	return repo.NewCompanyRepository(db)
}

func (c Container) GetSessionRepository(db mysql.DBManager) repository.SessionRepository {
	return repo.NewSessionRepository(db)
}

func (c Container) GetPlanRepository(db mysql.DBManager) repository.PlanRepository {
	return repo.NewPlanRepository(db)
}

func (c Container) GetRecoverSessionRepository(db mysql.DBManager) repository.RecoverSessionRepository {
	return repo.NewRecoverSessionRepository(db)
}
