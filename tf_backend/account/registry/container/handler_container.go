package container

import (
	"github.com/TrendFindProject/tf_backend/account/application/service"
	"github.com/TrendFindProject/tf_backend/account/interfaces/handler"
	pbAccount "github.com/TrendFindProject/tf_backend/proto/account/go"
)

func (c Container) GetAccountService(su service.UserService) pbAccount.UserServiceServer {
	return handler.NewAccountHandler(su)
}
