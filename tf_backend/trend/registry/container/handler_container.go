package container

import (
	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
	"github.com/TrendFindProject/tf_backend/trend/application/service"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/handler"
)

func (c Container) GetTrendServiceServer(ts service.TrendService) pbTrend.TrendServiceServer {
	return handler.NewTrendHandler(ts)
}
