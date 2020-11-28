package mysql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager interface {
	SQLManager
	Begin() (TxManager, error)
}

// NewDBManager generates and returns DBManager.
func NewDBManager(conn *sqlx.DB) DBManager {
	return &dbManager{conn}
}

// dbManager is the manager of SQL.
type dbManager struct {
	Conn *sqlx.DB
}

// SQLManager implement
// ExecContext executes SQL with context on sqlx.
func (s *dbManager) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return s.Conn.NamedExecContext(ctx, query, arg)
}

// SQLManager implement
// QueryContext executes query which return row with context on sqlx.
func (s *dbManager) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return s.Conn.NamedQueryContext(ctx, query, arg)
}

// Begin begins tx on sqlx.
func (s *dbManager) Begin() (TxManager, error) {
	tx, err := s.Conn.Beginx()
	if err != nil {
		return nil, err
	}
	return &txManager{tx}, nil
}
