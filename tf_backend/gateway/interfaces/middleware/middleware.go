package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/TrendFindProject/tf_backend/gateway/constant"
	redis "github.com/TrendFindProject/tf_backend/gateway/infrastructure/redis/client"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/support"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func Tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get trace id from request header
		traceID := r.Header.Get(constant.XRequestIDKey)
		if traceID == "" {
			traceID = xid.New().String()
		}

		// set trace id
		ctx := support.AddTraceIDToContext(r.Context(), traceID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logging
		ctx := r.Context()

		fields := []zap.Field{
			zap.String("Method", r.Method),
			zap.String("Request", r.RequestURI),
			zap.String("TraceID", support.GetTraceIDFromContext(ctx)),
		}

		// account api doesn't log user uuid
		if !strings.Contains(r.RequestURI, "account") {
			fields = append(fields, zap.String("UserUUID", fmt.Sprintf("%s", support.GetUserFromContext(ctx))))
		}

		logger.Common.Info(
			"GatewayConnectionLog",
			fields...,
		)

		next.ServeHTTP(w, r)
	})
}

func NewAuthentication() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				// FIXME: redisクローズは必要なのか...？
				//if err := redis.Client.Close(); err != nil {
				//	logger.Common.Info(err.Error())
				//	http.Error(w, InValidCookie, http.StatusForbidden)
				//	return
				//}
			}()

			// get cookie from request
			c, err := r.Cookie("VTKT")
			if err != nil {
				logger.Common.Warn(err.Error())
				w.WriteHeader(http.StatusForbidden)
				return
			}
			vtkt := c.Value

			// get user id from session
			uid, err := redis.TokenClient.HGet(vtkt, "uid").Result()
			if uid == "" || err == constant.RedisEmpty {
				logger.Common.Warn(err.Error())
				w.WriteHeader(http.StatusForbidden)
				return
			}

			// get company id from session
			cid, err := redis.TokenClient.HGet(vtkt, "cid").Result()
			if cid == "" || err == constant.RedisEmpty {
				logger.Common.Warn(err.Error())
				w.WriteHeader(http.StatusForbidden)
				return
			}

			// trace user id
			ctx := support.AddUserToContext(r.Context(), uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func NewCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if os.Getenv("ENV") == "LOCAL" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:30080")
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "https://trend-find.work")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}
