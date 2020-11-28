package repository

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
)

type CompanyRepository interface {
	FindCompanyById(context.Context, uint64) (*model.Company, error)
}
