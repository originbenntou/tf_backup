package main

import (
	"context"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TrendFindProject/tf_backend/gateway/constant"
	"github.com/TrendFindProject/tf_backend/gateway/graphql/account"
	accountGen "github.com/TrendFindProject/tf_backend/gateway/graphql/account/generated"
	"github.com/TrendFindProject/tf_backend/gateway/graphql/trend"
	trendGen "github.com/TrendFindProject/tf_backend/gateway/graphql/trend/generated"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/middleware"
	"github.com/gorilla/mux"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func init() {
	loc, err := time.LoadLocation(constant.Location)
	if err != nil {
		loc = time.FixedZone(constant.Location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = constant.DefaultPort
	}

	r := mux.NewRouter()
	r.Use(middleware.Tracing)
	// playground接続時は定期的にログが記録される...
	r.Use(middleware.Logging)
	r.Use(middleware.NewCORS)

	auth := middleware.NewAuthentication()

	aSvc := handler.NewDefaultServer(accountGen.NewExecutableSchema(accountGen.Config{Resolvers: account.NewAccountResolver()}))
	tSvc := handler.NewDefaultServer(trendGen.NewExecutableSchema(trendGen.Config{Resolvers: trend.NewTrendResolver()}))
	// custom error
	aSvc.SetErrorPresenter(
		func(ctx context.Context, err error) *gqlerror.Error {
			return graphql.DefaultErrorPresenter(ctx, &gqlerror.Error{
				Message:    err.Error(),
				Extensions: map[string]interface{}{"code": status.Code(err)},
			})
		},
	)
	tSvc.SetErrorPresenter(
		func(ctx context.Context, err error) *gqlerror.Error {
			return graphql.DefaultErrorPresenter(ctx, &gqlerror.Error{
				Message:    err.Error(),
				Extensions: map[string]interface{}{"code": status.Code(err)},
			})
		},
	)

	r.Path("/account").Handler(aSvc)
	r.Path("/trend").Handler(auth(tSvc))

	// HealthCheck
	r.Path("/health").HandlerFunc(healthCheckHandler)

	// GraphQL playground for local
	if os.Getenv("ENV") == "LOCAL" {
		r.Path("/").HandlerFunc(playground.Handler("GraphQL playground", "/account"))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Common.Fatal(err.Error())
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
