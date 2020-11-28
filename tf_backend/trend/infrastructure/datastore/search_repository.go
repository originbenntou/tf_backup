package datastore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TrendFindProject/tf_backend/trend/constant"
	"github.com/TrendFindProject/tf_backend/trend/domain/model"
	"github.com/TrendFindProject/tf_backend/trend/domain/repository"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type searchRepository struct {
	db mysql.DBManager
}

func NewSearchRepository(db mysql.DBManager) repository.SearchRepository {
	return &searchRepository{db}
}

func (r searchRepository) FindSearchBySearchWord(ctx context.Context, sw string, period time.Time) (m *model.Search, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM search WHERE search_word = :search_word AND created_at > :period_start"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"search_word": sw, "period_start": period.Format("2006-01-02 15:04:05")})
	if err != nil {
		return nil, err
	}

	m = &model.Search{}
	var list []*model.Search
	for rows.Next() {
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return nil, nil
	}

	// more than one record is error
	if len(list) > 1 {
		return nil, errors.New("found search more than 1 by: " + sw)
	}

	// one match record
	return list[0], nil
}

func (r searchRepository) CreateSearch(ctx context.Context, req *model.Search) (searchId uint64, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return constant.InvalidID, err
	}

	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}

		if txErr := tx.CloseTransaction(err); txErr != nil {
			logger.Common.Error(txErr.Error())
		}
	}()

	q := "INSERT INTO search (search_word, date, status, created_at, updated_at) VALUES (:search_word, :date, :status, :created_at, :updated_at)"

	result, err := tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return constant.InvalidID, err
	}

	if affect != 1 {
		msg := fmt.Sprintf("total affected: %d", affect)
		return constant.InvalidID, errors.New(msg)
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return constant.InvalidID, err
	}

	return uint64(lid), nil
}

func (r searchRepository) FindSearchById(ctx context.Context, sid uint64) (s *model.Search, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := `SELECT * FROM search WHERE id = :search_id`

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"search_id": sid})
	if err != nil {
		return nil, err
	}

	var list []*model.Search
	for rows.Next() {
		m := new(model.Search)
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return nil, nil
	}

	// more than one record is error
	if len(list) > 1 {
		return nil, errors.New("found search more than 1 by: " + fmt.Sprintf("%d", sid))
	}

	// one match record
	return list[0], nil
}
