package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	pbTrend "github.com/TrendFindProject/tf_backend/proto/trend/go"
	"github.com/TrendFindProject/tf_backend/trend/domain/model"
	"github.com/TrendFindProject/tf_backend/trend/domain/repository"
	"github.com/TrendFindProject/tf_backend/trend/util/pubsub"
)

type TrendService interface {
	TrendSearch(context.Context, *pbTrend.TrendSearchRequest) (*pbTrend.TrendSearchResponse, error)
	TrendSuggest(context.Context, *pbTrend.TrendSuggestRequest) (*pbTrend.TrendSuggestResponse, error)
	TrendHistory(context.Context, *pbTrend.TrendHistoryRequest) (*pbTrend.TrendHistoryResponse, error)
}

type trendService struct {
	repository.SearchRepository
	repository.SuggestRepository
	repository.HistoryRepository
}

func NewTrendService(
	sr repository.SearchRepository,
	sgr repository.SuggestRepository,
	hr repository.HistoryRepository,
) TrendService {
	return &trendService{
		sr,
		sgr,
		hr,
	}
}

func (s trendService) TrendSearch(ctx context.Context, pbReq *pbTrend.TrendSearchRequest) (*pbTrend.TrendSearchResponse, error) {
	var searchId uint64

	p := getPeriod()

	todaySearch, err := s.SearchRepository.FindSearchBySearchWord(ctx, pbReq.SearchWord, p)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// if keyword was not searched today, create new suggest and ML start
	if todaySearch == nil {
		searchId, err = s.SearchRepository.CreateSearch(ctx, &model.Search{
			SearchWord: pbReq.GetSearchWord(),
			Date:       p.Format("2006-01-02"),
			Status:     0,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// send message to ML
		if err = pubsub.PublishSingleGoroutine(ctx, pbReq.SearchWord, searchId); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		searchId = todaySearch.Id
	}

	todayHistory, err := s.HistoryRepository.FindHistoryByUserUuid(ctx, pbReq.UserUuid, searchId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// if history does not exist, create new history
	if todayHistory == nil {
		if err = s.HistoryRepository.CreateHistory(ctx, &model.History{
			UserUuid:  pbReq.UserUuid,
			SearchId:  searchId,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pbTrend.TrendSearchResponse{
		SearchId: searchId,
	}, nil
}

func (s trendService) TrendSuggest(ctx context.Context, pbReq *pbTrend.TrendSuggestRequest) (*pbTrend.TrendSuggestResponse, error) {
	suggests, err := s.SuggestRepository.FindSuggestsBySearchId(ctx, pbReq.SearchId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if suggests == nil {
		return nil, status.Error(codes.NotFound, "suggest not found, search id: "+fmt.Sprintf("%d", pbReq.SearchId))
	}

	// mapping for pb response
	var ss []*pbTrend.Suggest
	for _, s := range suggests {
		ok, i := getAssignWordIndex(ss, s.SuggestWord)
		if !ok {
			ss = append(ss, &pbTrend.Suggest{
				SuggestWord: s.SuggestWord,
			})
		}

		if ss == nil {
			return nil, status.Error(codes.Internal, "invalid suggest list")
		}

		var sg, mg, lg []*pbTrend.Graph
		err = json.Unmarshal([]byte(s.ChildSuggest.ShortGraphs), &sg)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		err = json.Unmarshal([]byte(s.ChildSuggest.MediumGraphs), &mg)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		err = json.Unmarshal([]byte(s.ChildSuggest.LongGraphs), &lg)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// suggest length is less than expected
		if len(ss) < i+1 {
			return nil, status.Error(codes.Internal, "invalid suggest list")
		}

		ss[i].ChildSuggests = append(ss[i].ChildSuggests, &pbTrend.ChildSuggest{
			ChildSuggestWord: s.ChildSuggest.ChildSuggestWord,
			Growth: &pbTrend.Growth{
				Short:  pbTrend.Arrow(s.ChildSuggest.Short),
				Medium: pbTrend.Arrow(s.ChildSuggest.Medium),
				Long:   pbTrend.Arrow(s.ChildSuggest.Long),
			},
			Graphs: &pbTrend.Graphs{
				Short:  sg,
				Medium: mg,
				Long:   lg,
			},
		})
	}

	return &pbTrend.TrendSuggestResponse{
		Suggest: ss,
	}, nil
}

func (s trendService) TrendHistory(ctx context.Context, pbReq *pbTrend.TrendHistoryRequest) (*pbTrend.TrendHistoryResponse, error) {
	histories, err := s.HistoryRepository.FindHistoryWithSuggestInfoByUserUuid(ctx, pbReq.UserUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if histories == nil {
		// correct even if no user history
		logger.Common.Info("user history not found, user_uuid: " + fmt.Sprintf("%d", pbReq.UserUuid))
		return nil, nil
	}

	var hs []*pbTrend.History
	for _, h := range histories {
		hs = append(hs, &pbTrend.History{
			SearchWord: h.SearchWord,
			SearchId:   h.SearchId,
			Date:       h.Date,
			Status:     pbTrend.Progress(h.Status),
		})
	}

	// max 5 items order by desc
	max := 5
	if len(hs) < 5 {
		max = len(hs)
	}

	return &pbTrend.TrendHistoryResponse{
		Histories: hs[:max],
	}, nil
}

// サジェスト情報は毎日4時切り替えのため、DB操作の起点時間を計算する
// 検索当日中に同じキーワードで検索されていれば同じサジェストを返却
func getPeriod() time.Time {
	n := time.Now()
	p := time.Date(n.Year(), n.Month(), n.Day(), 4, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

	// 検索時間が4時より前なら前日の4時を起点とする
	if n.Before(p) {
		p = p.AddDate(0, 0, -1)
	}

	return p
}

// サジェストワードを割り当てる配列のインデックスを求める
// サジェストワードが含まれていればそのtrueとインデックスを、含まれていなければfalseと新たなインデックスを返却
func getAssignWordIndex(ss []*pbTrend.Suggest, w string) (bool, int) {
	for i, s := range ss {
		if w == s.SuggestWord {
			return true, i
		}
	}

	return false, len(ss)
}
