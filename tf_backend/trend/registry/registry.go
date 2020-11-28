package registry

import (
	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/mysql"
	"github.com/TrendFindProject/tf_backend/trend/registry/container"
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
	pbTrend.RegisterTrendServiceServer(r.Server,
		r.GetTrendServiceServer(
			r.GetTrendService(
				r.GetSearchRepository(r.DBManager),
				r.GetSuggestRepository(r.DBManager),
				r.GetHistoryRepository(r.DBManager),
			),
		),
	)

	//grpchealth?
}
