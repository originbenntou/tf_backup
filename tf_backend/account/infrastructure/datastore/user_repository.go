package datastore

import (
	"context"
	"errors"
	"fmt"

	"github.com/TrendFindProject/tf_backend/account/constant"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	"github.com/TrendFindProject/tf_backend/account/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db mysql.DBManager
}

func NewUserRepository(db mysql.DBManager) repository.UserRepository {
	return &userRepository{db}
}

func (r userRepository) FindUserByEmail(ctx context.Context, email string) (u *model.User, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM user WHERE email = :email"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"email": email})
	if err != nil {
		return nil, err
	}

	var list []*model.User
	for rows.Next() {
		m := new(model.User)
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
		return nil, errors.New("found user more than 1 by: " + email)
	}

	// one match record
	return list[0], nil
}

func (r userRepository) FindUserByUuid(ctx context.Context, uid string) (u *model.User, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM user WHERE uuid = :uid"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"uid": uid})
	if err != nil {
		return nil, err
	}

	var list []*model.User
	for rows.Next() {
		m := new(model.User)
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
		return nil, errors.New("found user more than 1 by: " + uid)
	}

	// one match record
	return list[0], nil
}

func (r userRepository) CreateUser(ctx context.Context, req *model.User) (id uint64, err error) {
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

	q := "INSERT INTO user (uuid, email, password, name, company_id, created_at, updated_at) VALUES (:uuid, :email, :password, :name, :company_id, :created_at, :updated_at)"

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

func (r userRepository) CountUsersByCompanyId(ctx context.Context, cid uint64) (c uint64, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT COUNT(*) FROM user WHERE company_id = :company_id"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"company_id": cid})
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		if err := rows.Scan(&c); err != nil {
			return 0, err
		}
	}

	return
}

func (r userRepository) CountUsersByUuid(ctx context.Context, uid string) (c uint64, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT COUNT(*) FROM user WHERE uuid = :uuid"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"uuid": uid})
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		if err := rows.Scan(&c); err != nil {
			return 0, err
		}
	}

	return
}

func (r userRepository) ModifyUserPassword(ctx context.Context, uid string, password []byte) (err error) {
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

	q := "UPDATE user SET password = :password WHERE uuid = :uuid"

	result, err := tx.NamedExecContext(ctx, q, map[string]interface{}{"password": password, "uuid": uid})
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

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
