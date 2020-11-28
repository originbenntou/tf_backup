package registry

import (
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
	"github.com/TrendFindProject/tf_backend/account/registry/container"
	pbAccount "github.com/TrendFindProject/tf_backend/proto/account/go"
	"google.golang.org/grpc"
)

type Registry interface {
	Register()
}

type registry struct {
	*grpc.Server
	mysql.DBManager
	container.Container //DI Container
}

func NewRegistry(s *grpc.Server, db mysql.DBManager) Registry {
	return &registry{s, db, container.Container{}}
}

func (r registry) Register() {
	pbAccount.RegisterUserServiceServer(r.Server,
		r.GetAccountService(
			r.GetUserService(
				r.GetUserRepository(r.DBManager),
				r.GetCompanyRepository(r.DBManager),
				r.GetSessionRepository(r.DBManager),
				r.GetPlanRepository(r.DBManager),
				r.GetRecoverSessionRepository(r.DBManager),
			),
		),
	)

	//grpchealth?
}
