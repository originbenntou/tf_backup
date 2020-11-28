package handler

import (
	"context"
	"github.com/TrendFindProject/tf_backend/account/application/service"

	pbAccount "github.com/TrendFindProject/tf_backend/proto/account/go"
)

type AccountHandler interface {
	RegisterUser(context.Context, *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error)
	VerifyUser(context.Context, *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error)
}

type accountHandler struct {
	service.UserService
}

func NewAccountHandler(us service.UserService) pbAccount.UserServiceServer {
	return &accountHandler{us}
}
