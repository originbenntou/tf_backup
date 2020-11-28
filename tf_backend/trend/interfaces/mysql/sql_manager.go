package mysql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

// SQLManager is the manager of DB.
type SQLManager interface {
	Querier
	//Preparer
	Executor
}

// 使ってないものは整理
// prepareはいつか使うかも知れない...
type (
	// Querier is interface of Query.
	Querier interface {
		//Query(query string, args ...interface{}) (*sql.Rows, error)
		//QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
		NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	}

	// Executor is interface of Execute.
	Executor interface {
		//Exec(query string, args ...interface{}) (sql.Result, error)
		//ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	}

	// Preparer is interface of Prepare.
	Preparer interface {
		//Prepare(query string) (*sql.Stmt, error)
		//PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	}
)
