package container

import (
	"github.com/TrendFindProject/tf_backend/account/application/service"
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
)

func (c Container) GetUserService(
	ur repository.UserRepository,
	cr repository.CompanyRepository,
	sr repository.SessionRepository,
	pr repository.PlanRepository,
	rsr repository.RecoverSessionRepository,
) service.UserService {
	return service.NewUserService(ur, cr, sr, pr, rsr)
}
