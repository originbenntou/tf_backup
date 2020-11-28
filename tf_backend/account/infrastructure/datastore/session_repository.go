package datastore

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TrendFindProject/tf_backend/account/constant"
	"github.com/TrendFindProject/tf_backend/account/domain/model"
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	"github.com/TrendFindProject/tf_backend/account/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type sessionRepository struct {
	db mysql.DBManager
}

func NewSessionRepository(db mysql.DBManager) repository.SessionRepository {
	return &sessionRepository{db}
}

func (r sessionRepository) FindExistTokenByUserUuid(ctx context.Context, uid string) (token string, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT token FROM session WHERE user_uuid = :user_uuid AND DATE_ADD(updated_at, INTERVAL " + constant.SessionExpireHour + ") > NOW()"

	rows, err := r.db.NamedQueryContext(ctx, q, map[string]interface{}{"user_uuid": uid})
	if err != nil {
		return "", err
	}

	var list []string
	for rows.Next() {
		t := new(string)
		if err := rows.Scan(t); err != nil {
			return "", err
		}
		list = append(list, *t)
	}

	// no match record is ok, return empty
	if len(list) == 0 {
		return "", nil
	}

	// more than one record is error
	if len(list) > 1 {
		return "", errors.New("found token more than 1 by: " + string(uid))
	}

	// one match record
	return list[0], nil
}

func (r sessionRepository) CreateSession(ctx context.Context, req *model.Session) (err error) {
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

	q := "INSERT INTO session (token, user_uuid, company_id) VALUES (:token, :user_uuid, :company_id)"

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

func (r sessionRepository) DeleteSessionByUserUuid(ctx context.Context, uid string) (err error) {
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

	q := "DELETE FROM session WHERE user_uuid = :user_uuid AND DATE_ADD(updated_at, INTERVAL " + constant.SessionExpireHour + ") > NOW()"

	result, err := tx.NamedExecContext(ctx, q, map[string]interface{}{
		"updated_at": time.Now().Format("2006-01-02 15:04:05.000000"), // micro second
		"user_uuid":  uid,
	})
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

func (r sessionRepository) CountValidSessionByCompanyId(ctx context.Context, cid uint64) (c uint64, err error) {
	defer func() {
		if err != nil {
			logger.Common.Error(err.Error())
		}
	}()

	q := "SELECT COUNT(*) FROM session WHERE company_id = :company_id AND DATE_ADD(updated_at, INTERVAL " + constant.SessionExpireHour + ") > NOW()"

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
