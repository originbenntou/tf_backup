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

type historyRepository struct {
	db mysql.DBManager
}

func NewHistoryRepository(db mysql.DBManager) repository.HistoryRepository {
	return &historyRepository{db}
}

func (r historyRepository) FindHistoryByUserUuid(ctx context.Context, uid string, sid uint64) (h *model.History, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := `SELECT * FROM history
			WHERE user_uuid = :user_uuid
			AND search_id = :search_id`

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"user_uuid": uid, "search_id": sid})
	if err != nil {
		return nil, err
	}

	var list []*model.History
	for rows.Next() {
		m := new(model.History)
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
		return nil, errors.New("found history more than 1 by: " + fmt.Sprintf("user_uuid=%s ", uid) + fmt.Sprintf("search_id=%d", sid))
	}

	// one match record
	return list[0], nil
}

func (r historyRepository) CreateHistory(ctx context.Context, req *model.History) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}

		if txErr := tx.CloseTransaction(err); txErr != nil {
			logger.Common.Error(txErr.Error())
		}
	}()

	q := "INSERT INTO history (user_uuid, search_id) VALUES (:user_uuid, :search_id)"

	result, err := tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		msg := fmt.Sprintf("total affected: %d", affect)
		return errors.New(msg)
	}

	return nil
}

func (r historyRepository) FindHistoryWithSuggestInfoByUserUuid(ctx context.Context, uid string) (hs []*model.History, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := `SELECT history.*, search.search_word, search.date, search.status
			FROM history 
			LEFT JOIN search ON history.search_id = search.id
			WHERE user_uuid = :user_uuid
			ORDER BY history.id DESC`

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"user_uuid": uid})
	if err != nil {
		return nil, err
	}

	var list []*model.History
	for rows.Next() {
		m := new(model.History)
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		list = append(list, m)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return nil, nil
	}

	return list, nil
}
