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

type recoverSessionRepository struct {
	db mysql.DBManager
}

func NewRecoverSessionRepository(db mysql.DBManager) repository.RecoverSessionRepository {
	return &recoverSessionRepository{db}
}

func (r recoverSessionRepository) CreateRecoverSession(ctx context.Context, req *model.RecoverSession) (err error) {
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

	// one recover_session for each user
	q := `INSERT INTO recover_session (user_uuid, auth_key, recover_token)
			VALUES (:user_uuid, :auth_key, :recover_token)
			ON DUPLICATE KEY UPDATE user_uuid = :user_uuid, auth_key = :auth_key, recover_token = :recover_token`

	result, err := tx.NamedExecContext(ctx, q, req)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// affect is 2 records for upsert, 1 record for insert
	if affect > 2 {
		msg := fmt.Sprintf("total affected: %d", affect)
		return errors.New(msg)
	}

	return nil
}

func (r recoverSessionRepository) FindRecoverSessionByUuid(ctx context.Context, uid string) (rs *model.RecoverSession, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT * FROM recover_session WHERE user_uuid = :user_uuid"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"user_uuid": uid})
	if err != nil {
		return nil, err
	}

	var list []*model.RecoverSession
	for rows.Next() {
		m := new(model.RecoverSession)
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
		return nil, errors.New("found recover_session more than 1 by: " + uid)
	}

	// one match record
	return list[0], nil
}

func (r recoverSessionRepository) DeleteRecoverSessionByUuid(ctx context.Context, uid string) (err error) {
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

	q := `DELETE FROM recover_session WHERE user_uuid = :user_uuid`

	result, err := tx.NamedExecContext(ctx, q, map[string]interface{}{"user_uuid": uid})
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
