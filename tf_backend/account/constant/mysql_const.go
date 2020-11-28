package constant

import (
	"os"
	"time"
)

const (
	MaxIdleConns    = 100
	MaxOpenConns    = 10
	ConnMaxLifetime = 0
	InvalidID       = 0
)

type accountConfig struct {
	host            string
	port            string
	user            string
	password        string
	dbname          string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
}

func (c accountConfig) GetHost() string                   { return c.host }
func (c accountConfig) GetPort() string                   { return c.port }
func (c accountConfig) GetUser() string                   { return c.user }
func (c accountConfig) GetPassword() string               { return c.password }
func (c accountConfig) GetDbname() string                 { return c.dbname }
func (c accountConfig) GetMaxIdleConns() int              { return c.maxIdleConns }
func (c accountConfig) GetMaxOpenConns() int              { return c.maxOpenConns }
func (c accountConfig) GetConnMaxLifetime() time.Duration { return c.connMaxLifetime }

var Config accountConfig

func init() {
	switch os.Getenv("ENV") {
	case "LOCAL":
		Config = accountConfig{
			host:            "mysql",
			port:            "3306",
			user:            "2929",
			password:        "2929",
			dbname:          "account",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	case "PRODUCTION":
		Config = accountConfig{
			host:            "127.0.0.1",
			port:            "3306",
			user:            "2929",
			password:        "dmeej1pu4Cos21gO",
			dbname:          "account",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	default:
		Config = accountConfig{
			host:            "mysql",
			port:            "3306",
			user:            "2929",
			password:        "2929",
			dbname:          "account",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	}
}
