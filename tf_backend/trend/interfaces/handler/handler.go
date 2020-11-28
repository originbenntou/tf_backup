package handler

import (
	"context"

	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
	"github.com/TrendFindProject/tf_backend/trend/application/service"
)

type TrendHandler interface {
	TrendSearch(context.Context, *pbTrend.TrendSearchRequest) (*pbTrend.TrendSearchResponse, error)
	TrendSuggest(context.Context, *pbTrend.TrendSuggestRequest) (*pbTrend.TrendSuggestResponse, error)
	TrendHistory(context.Context, *pbTrend.TrendHistoryRequest) (*pbTrend.TrendHistoryResponse, error)
}

type trendHandler struct {
	service.TrendService
}

func NewTrendHandler(ts service.TrendService) pbTrend.TrendServiceServer {
	return &trendHandler{ts}
}
