package datastore

import (
	"context"
	"errors"

	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	"github.com/TrendFindProject/tf_backend/account/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type planRepository struct {
	db mysql.DBManager
}

func NewPlanRepository(db mysql.DBManager) repository.PlanRepository {
	return &planRepository{db}
}

func (r planRepository) FindCapacityById(ctx context.Context, id uint64) (cap uint64, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT capacity FROM plan WHERE id = :id"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"id": id})
	if err != nil {
		return 0, err
	}

	var list []uint64
	for rows.Next() {
		c := new(uint64)
		if err := rows.Scan(c); err != nil {
			return 0, err
		}
		list = append(list, *c)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return 0, nil
	}

	// more than one record is error
	if len(list) > 1 {
		return 0, errors.New("found capacity more than 1 by plan_id: " + string(id))
	}

	// one match record
	return list[0], nil
}
