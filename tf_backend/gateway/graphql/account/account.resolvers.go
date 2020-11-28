package account

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	"github.com/TrendFindProject/tf_backend/gateway/constant"
	"github.com/TrendFindProject/tf_backend/gateway/graphql/account/generated"
	"github.com/TrendFindProject/tf_backend/gateway/graphql/account/model"
	redis "github.com/TrendFindProject/tf_backend/gateway/infrastructure/redis/client"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/logger"
	"github.com/TrendFindProject/tf_backend/gateway/interfaces/support"
	pbAccount "github.com/TrendFindProject/tf_backend/proto/account/go"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, user model.User) (ok bool, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	pbUser, err := r.accountClient.RegisterUser(ctx, &pbAccount.RegisterUserRequest{
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		CompanyId: uint64(user.CompanyID),
	})
	if err != nil {
		return false, err
	}

	if pbUser == nil || pbUser.UserUuid == "" {
		return false, status.Error(codes.Internal, "incorrect gRPC response")
	}

	return true, nil
}

func (r *queryResolver) VerifyUser(ctx context.Context, email string, password string) (newToken string, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	pbRes, err := r.accountClient.VerifyUser(ctx, &pbAccount.VerifyUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	newToken = pbRes.NewToken
	oldToken := pbRes.OldToken
	uid := pbRes.User.Uuid
	cid := pbRes.User.CompanyId

	// delete old token if double login
	if oldToken != "" {
		if err = redis.TokenClient.Del(oldToken).Err(); err != nil {
			return "", err
		}
	}

	// set token with user_uuid and company_id to Redis
	if err = redis.TokenClient.HMSet(newToken, map[string]interface{}{
		"uid": uid,
		"cid": cid,
	}).Err(); err != nil {
		return "", err
	}
	// expire 6 days
	if err = redis.TokenClient.Expire(newToken, time.Hour*constant.SessionExpireHour).Err(); err != nil {
		return "", err
	}

	return newToken, nil
}

func (r *queryResolver) SendRecoverEmail(ctx context.Context, email string, name string) (authKey string, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	pbRes, err := r.accountClient.SendRecoverEmail(ctx, &pbAccount.SendRecoverEmailRequest{
		Email: email,
		Name:  name,
	})
	if err != nil {
		return "", err
	}

	return pbRes.AuthKey, nil
}

func (r *mutationResolver) RecoverPassword(ctx context.Context, token string, authKey string, password string) (ok bool, err error) {
	defer func() {
		if err != nil {
			logger.Common.Info(err.Error(), zap.String("TraceID", support.GetTraceIDFromContext(ctx)))
		}
	}()

	_, err = r.accountClient.RecoverPassword(ctx, &pbAccount.RecoverPasswordRequest{
		RecoverToken: token,
		AuthKey:      authKey,
		Password:     password,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ModifyPassword(ctx context.Context, oldPassword string, newPassword string) (bool, error) {
	panic("implement me")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
