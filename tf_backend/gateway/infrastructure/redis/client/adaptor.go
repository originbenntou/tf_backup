package client

import (
	"os"

	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	goRedis "github.com/go-redis/redis/v7"
)

// ログインチェック
// Hash型 key: Token, field: uid,cid
var TokenClient *goRedis.Client

// FIXME: configをつかって書き直す
func init() {
	TokenClient = goRedis.NewClient(&goRedis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Redis疎通確認
	var err error

	_, err = TokenClient.Ping().Result()
	if err != nil {
		logger.Common.Error(err.Error())
	}
}
