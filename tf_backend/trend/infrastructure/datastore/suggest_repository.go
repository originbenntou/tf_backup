package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrendFindProject/tf_backend/trend/domain/model"
	"github.com/TrendFindProject/tf_backend/trend/domain/repository"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type suggestRepository struct {
	db mysql.DBManager
}

func NewSuggestRepository(db mysql.DBManager) repository.SuggestRepository {
	return &suggestRepository{db}
}

func (r suggestRepository) FindSuggestsBySearchId(ctx context.Context, sid uint64) (cs []*model.Suggest, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := `SELECT suggest.suggest_word, child_suggest.*
			FROM suggest 
			LEFT JOIN child_suggest ON suggest.id = child_suggest.suggest_id
			WHERE search_id = :search_id`

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"search_id": sid})
	if err != nil {
		return nil, err
	}

	var list []*model.Suggest
	for rows.Next() {
		m := new(model.Suggest)
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return nil, nil
	}

	// more than six record is error
	if len(list) > 36 {
		return nil, errors.New("found history more than 36 by: " + fmt.Sprintf("search_id=%d", sid))
	}

	return list, nil
}
