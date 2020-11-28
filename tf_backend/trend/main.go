package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TrendFindProject/tf_backend/trend/constant"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/interceptor"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/trend/interfaces/mysql"
	"github.com/TrendFindProject/tf_backend/trend/registry"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			interceptor.XTraceID(),
			interceptor.XUserUUID(),
			grpcZap.UnaryServerInterceptor(logger.Interceptor),
			interceptor.Logging(),
		)),
	)

	conn, err := mysql.NewDBConnection(constant.Config)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}
	db := mysql.NewDBManager(conn)

	registry.NewRegistry(srv, db).Register()
	// server reflection
	reflection.Register(srv)

	go func() {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to create listener: %s", err)
		}
		log.Println("start server on port", port)
		if err := srv.Serve(listener); err != nil {
			log.Println("failed to exit serve: ", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint
	log.Println("received a signal of graceful shutdown")

	stopped := make(chan struct{})
	go func() {
		srv.GracefulStop()
		close(stopped)
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(), 1*time.Minute)

	select {
	case <-ctx.Done():
		srv.Stop()
	case <-stopped:
		cancel()
	}

	log.Println("completed graceful shutdown")
}
