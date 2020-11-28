package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDbname() string
	GetMaxIdleConns() int
	GetMaxOpenConns() int
	GetConnMaxLifetime() time.Duration
}

func NewDBConnection(c Config) (*sqlx.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), c.GetDbname())
	db, err := sqlx.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(c.GetMaxIdleConns())
	db.SetMaxOpenConns(c.GetMaxOpenConns())
	db.SetConnMaxLifetime(c.GetConnMaxLifetime())

	return db, nil
}
