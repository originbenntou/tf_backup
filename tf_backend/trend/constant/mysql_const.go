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

type trendConfig struct {
	host            string
	port            string
	user            string
	password        string
	dbname          string
	maxIdleConns    int
	maxOpenConns    int
	connMaxLifetime time.Duration
}

func (c trendConfig) GetHost() string                   { return c.host }
func (c trendConfig) GetPort() string                   { return c.port }
func (c trendConfig) GetUser() string                   { return c.user }
func (c trendConfig) GetPassword() string               { return c.password }
func (c trendConfig) GetDbname() string                 { return c.dbname }
func (c trendConfig) GetMaxIdleConns() int              { return c.maxIdleConns }
func (c trendConfig) GetMaxOpenConns() int              { return c.maxOpenConns }
func (c trendConfig) GetConnMaxLifetime() time.Duration { return c.connMaxLifetime }

var Config trendConfig

func init() {
	switch os.Getenv("ENV") {
	case "LOCAL":
		Config = trendConfig{
			host:            "mysql",
			port:            "3306",
			user:            "2929",
			password:        "2929",
			dbname:          "trend",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	case "PRODUCTION":
		Config = trendConfig{
			host:            "127.0.0.1",
			port:            "3306",
			user:            "2929",
			password:        "dmeej1pu4Cos21gO",
			dbname:          "trend",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	default:
		Config = trendConfig{
			host:            "mysql",
			port:            "3306",
			user:            "2929",
			password:        "2929",
			dbname:          "trend",
			maxIdleConns:    MaxIdleConns,
			maxOpenConns:    MaxOpenConns,
			connMaxLifetime: ConnMaxLifetime,
		}
	}
}
