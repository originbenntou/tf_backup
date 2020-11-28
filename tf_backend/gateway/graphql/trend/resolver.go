package trend

import (
	"github.com/TrendFindProject/tf_backend/gateway/infrastructure/grpc/client"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/interceptor"
	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
	"os"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	trendClient pbTrend.TrendServiceClient
}

func NewTrendResolver() *Resolver {
	return &Resolver{
		trendClient: pbTrend.NewTrendServiceClient(
			client.GetGrpcConn(os.Getenv("TREND_ADDR"), interceptor.XTraceID, interceptor.XUserUUID),
		),
	}
}
