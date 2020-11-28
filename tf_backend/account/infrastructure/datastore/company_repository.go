package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrendFindProject/tf_backend/account/domain/model"
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	"github.com/TrendFindProject/tf_backend/account/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type companyRepository struct {
	db mysql.DBManager
}

func NewCompanyRepository(db mysql.DBManager) repository.CompanyRepository {
	return &companyRepository{db}
}

func (r companyRepository) FindCompanyById(ctx context.Context, id uint64) (m *model.Company, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM company WHERE id = :id"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	m = &model.Company{}
	c := 0
	for rows.Next() {
		if err := rows.StructScan(m); err != nil {
			return nil, err
		}
		c++
	}

	// no match record is ok, return empty
	if c == 0 {
		return nil, nil
	}

	// more than one record is error
	if c > 1 {
		msg := fmt.Sprintf("found company more than 1 by company_id: %d", id)
		return nil, errors.New(msg)
	}

	// one match record
	return m, nil
}
