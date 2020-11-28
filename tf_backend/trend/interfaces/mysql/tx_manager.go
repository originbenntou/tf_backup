package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TxManager is the manager of Tx.
type TxManager interface {
	Executor
	Commit() error
	Rollback() error
	CloseTransaction(err error) error
}

// txManager is the manager of Tx.
type txManager struct {
	*sqlx.Tx
}

// Executor implement
// ExecContext executes SQL with context on sqlx.
func (s *txManager) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return s.Tx.NamedExecContext(ctx, query, arg)
}

// Tx Commit
func (s *txManager) Commit() error {
	return s.Tx.Commit()
}

// Tx Rollback
func (s *txManager) Rollback() error {
	return s.Tx.Rollback()
}

// Close Tx
// Don't write panic() not to stop server
func (s *txManager) CloseTransaction(err error) error {
	if recover() != nil {
		return s.Rollback()
	} else if err != nil {
		return s.Rollback()
	} else {
		return s.Commit()
	}
}
