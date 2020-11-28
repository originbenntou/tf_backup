package trend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TrendFindProject/tf_backend/gateway/graphql/trend/generated"
	"github.com/TrendFindProject/tf_backend/gateway/graphql/trend/model"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/support"
	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
)

func (r *queryResolver) TrendSearch(ctx context.Context, keyword string) (suggestId int, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(),
				zap.String("UserUUID", fmt.Sprintf("%s", support.GetUserFromContext(ctx))),
				zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	uid := support.GetUserFromContext(ctx)
	if uid == "" {
		return 0, status.Error(codes.Internal, "uuid not found in context")
	}

	pbRes, err := r.trendClient.TrendSearch(ctx, &pbTrend.TrendSearchRequest{
		SearchWord: keyword,
		UserUuid:   uid,
	})

	if err != nil {
		return 0, err
	}

	return int(pbRes.SearchId), nil
}

func (r *queryResolver) TrendSuggest(ctx context.Context, suggestID int) (suggests []*model.Suggest, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(),
				zap.String("UserUUID", fmt.Sprintf("%s", support.GetUserFromContext(ctx))),
				zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	pbRes, err := r.trendClient.TrendSuggest(ctx, &pbTrend.TrendSuggestRequest{
		SearchId: uint64(suggestID),
	})
	if err != nil {
		return nil, err
	}

	// mapping
	for _, s := range pbRes.Suggest {
		var childSuggests []*model.ChildSuggest
		for _, cs := range s.ChildSuggests {
			var sgs, mgs, lgs []*model.Graph
			for _, sg := range cs.Graphs.Short {
				sgs = append(sgs, &model.Graph{
					Date:  sg.Date,
					Value: int(sg.Value),
				})
			}

			for _, mg := range cs.Graphs.Medium {
				mgs = append(mgs, &model.Graph{
					Date:  mg.Date,
					Value: int(mg.Value),
				})
			}

			for _, lg := range cs.Graphs.Long {
				lgs = append(lgs, &model.Graph{
					Date:  lg.Date,
					Value: int(lg.Value),
				})
			}

			childSuggests = append(childSuggests, &model.ChildSuggest{
				Word: cs.ChildSuggestWord,
				Growth: &model.Growth{
					Short:  model.Arrow(cs.Growth.Short.String()),
					Medium: model.Arrow(cs.Growth.Medium.String()),
					Long:   model.Arrow(cs.Growth.Long.String()),
				},
				Graphs: &model.Graphs{
					Short:  sgs,
					Medium: mgs,
					Long:   lgs,
				},
			})
		}

		suggests = append(suggests, &model.Suggest{
			Keyword:       s.SuggestWord,
			ChildSuggests: childSuggests,
		})
	}

	return suggests, nil
}

func (r *queryResolver) TrendHistory(ctx context.Context) (histories []*model.History, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(),
				zap.String("UserUUID", fmt.Sprintf("%s", support.GetUserFromContext(ctx))),
				zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	uid := support.GetUserFromContext(ctx)
	if uid == "" {
		return nil, status.Error(codes.Internal, "uuid not found in context")
	}

	pbRes, err := r.trendClient.TrendHistory(ctx, &pbTrend.TrendHistoryRequest{
		UserUuid: uid,
	})
	if err != nil {
		return nil, err
	}

	for _, h := range pbRes.Histories {
		histories = append(histories, &model.History{
			SuggestID: int(h.SearchId),
			Keyword:   h.SearchWord,
			Date:      h.Date,
			Status:    model.Progress(h.Status.String()),
		})
	}

	return histories, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
