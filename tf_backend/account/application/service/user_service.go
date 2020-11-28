package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/TrendFindProject/tf_backend/account/interfaces/mail"
	"github.com/golang/protobuf/ptypes/empty"
	"time"

	"github.com/TrendFindProject/tf_backend/account/domain/model"
	"github.com/TrendFindProject/tf_backend/account/domain/repository"
	"github.com/TrendFindProject/tf_backend/account/interfaces/auth"
	pbAccount "github.com/TrendFindProject/tf_backend/proto/account/go"
	"github.com/google/uuid"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	RegisterUser(context.Context, *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error)
	VerifyUser(context.Context, *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error)
	SendRecoverEmail(context.Context, *pbAccount.SendRecoverEmailRequest) (*pbAccount.SendRecoverEmailResponse, error)
	RecoverPassword(context.Context, *pbAccount.RecoverPasswordRequest) (*empty.Empty, error)
}

type userService struct {
	repository.UserRepository
	repository.CompanyRepository
	repository.SessionRepository
	repository.PlanRepository
	repository.RecoverSessionRepository
}

func NewUserService(
	ur repository.UserRepository,
	cr repository.CompanyRepository,
	sr repository.SessionRepository,
	pr repository.PlanRepository,
	rsr repository.RecoverSessionRepository,
) UserService {
	return &userService{
		ur,
		cr,
		sr,
		pr,
		rsr,
	}
}

func (s userService) RegisterUser(ctx context.Context, pbReq *pbAccount.RegisterUserRequest) (*pbAccount.RegisterUserResponse, error) {
	user, err := s.UserRepository.FindUserByEmail(ctx, pbReq.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// if user exists
	if user != nil {
		return nil, status.Error(codes.AlreadyExists, errors.New("user already exist: "+pbReq.GetEmail()).Error())
	}

	company, err := s.CompanyRepository.FindCompanyById(ctx, pbReq.GetCompanyId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// if company not exists
	if company == nil {
		msg := fmt.Sprintf("company not found: %s", pbReq.GetEmail())
		return nil, status.Error(codes.InvalidArgument, errors.New(msg).Error())
	}

	capacity, err := s.PlanRepository.FindCapacityById(ctx, company.PlanId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	existUserCount, err := s.UserRepository.CountUsersByCompanyId(ctx, company.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// when new user create over plan capacity, forbidden
	if existUserCount+1 > capacity {
		msg := fmt.Sprintf("forbidden create user over plan capacity: %s", pbReq.GetEmail())
		return nil, status.Error(codes.OutOfRange, errors.New(msg).Error())
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(pbReq.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userUuid := xid.New().String()
	count, err := s.UserRepository.CountUsersByUuid(ctx, userUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// retry generate uuid once
	if count > 0 {
		userUuid = xid.New().String()
	}

	_, err = s.UserRepository.CreateUser(ctx, &model.User{
		Uuid:      userUuid,
		Email:     pbReq.GetEmail(),
		PassHash:  passHash,
		Name:      pbReq.GetName(),
		CompanyId: pbReq.GetCompanyId(),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbAccount.RegisterUserResponse{
		UserUuid: userUuid,
	}, nil
}

func (s userService) VerifyUser(ctx context.Context, pbReq *pbAccount.VerifyUserRequest) (*pbAccount.VerifyUserResponse, error) {
	user, err := s.UserRepository.FindUserByEmail(ctx, pbReq.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, errors.New("user is not found: "+pbReq.GetEmail()).Error())
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(pbReq.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// forbidden login over limit by plan
	isOver, err := s.isLoginOverCapacityOfCompanyPlan(ctx, user.CompanyId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if isOver {
		return nil, status.Error(codes.PermissionDenied, errors.New("forbidden login over limit by plan").Error())
	}

	// generate new token
	// if double login, delete old token
	newToken, oldToken, err := s.getCorrectTokenForUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbAccount.VerifyUserResponse{
		NewToken: newToken,
		OldToken: oldToken,
		User: &pbAccount.User{
			Id:        user.Id,
			Uuid:      user.Uuid,
			Email:     user.Email,
			Name:      user.Name,
			CompanyId: user.CompanyId,
		},
	}, nil
}

func (s userService) SendRecoverEmail(ctx context.Context, pbReq *pbAccount.SendRecoverEmailRequest) (*pbAccount.SendRecoverEmailResponse, error) {
	user, err := s.UserRepository.FindUserByEmail(ctx, pbReq.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, errors.New("user is not found: "+pbReq.GetEmail()).Error())
	}

	if user.Name != pbReq.GetName() {
		return nil, status.Error(codes.NotFound, errors.New("name is not match: "+pbReq.GetName()).Error())
	}

	uid := user.Uuid

	authKey := auth.GeneratePassword()
	if authKey == "" {
		return nil, status.Error(codes.Internal, "failed to generate auth key")
	}

	authHash, err := bcrypt.GenerateFromPassword([]byte(authKey), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// tokenString expire 24 hour
	recoverToken := auth.GenerateNewTokenByUuid(uid)

	if err := s.RecoverSessionRepository.CreateRecoverSession(ctx, &model.RecoverSession{
		UserUuid:     uid,
		AuthKeyHash:  authHash,
		RecoverToken: recoverToken,
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := mail.SendRecoverUrl(user.Email, user.Name, recoverToken); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbAccount.SendRecoverEmailResponse{
		AuthKey: authKey,
	}, nil
}

func (s userService) RecoverPassword(ctx context.Context, pbReq *pbAccount.RecoverPasswordRequest) (*empty.Empty, error) {
	// check valid recover token, expire,signature...etc!
	t, err := auth.ValidateTokenString(pbReq.RecoverToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !t.Valid {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid recover token: %v", pbReq.RecoverToken))
	}

	uid := auth.GetUserUuidFromClaim(t)
	if uid == "" {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed parse claim: %v", pbReq.RecoverToken))
	}

	rs, err := s.RecoverSessionRepository.FindRecoverSessionByUuid(ctx, uid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// verify recover token
	if rs.RecoverToken != pbReq.RecoverToken {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("request recover token not match: %v", pbReq.RecoverToken))
	}

	// verify auth key
	if err := bcrypt.CompareHashAndPassword(rs.AuthKeyHash, []byte(pbReq.AuthKey)); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// generate new password
	passHash, err := bcrypt.GenerateFromPassword([]byte(pbReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.UserRepository.ModifyUserPassword(ctx, uid, passHash); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// get user's name and email
	user, err := s.UserRepository.FindUserByUuid(ctx, rs.UserUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("invalid recover api call %v", rs.UserUuid))
	}

	// send complete mail
	if err = mail.SendRecoverComplete(user.Email, user.Name); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// because only use recover session once, delete
	// if error occur, try again from the beginning
	if err := s.RecoverSessionRepository.DeleteRecoverSessionByUuid(ctx, uid); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s userService) isLoginOverCapacityOfCompanyPlan(ctx context.Context, cid uint64) (bool, error) {
	count, err := s.SessionRepository.CountValidSessionByCompanyId(ctx, cid)
	if err != nil {
		return false, err
	}

	company, err := s.CompanyRepository.FindCompanyById(ctx, cid)
	if err != nil {
		return false, err
	}

	capacity, err := s.PlanRepository.FindCapacityById(ctx, company.PlanId)
	if err != nil {
		return false, err
	}

	return count > capacity, nil
}

func (s userService) getCorrectTokenForUser(ctx context.Context, user *model.User) (string, string, error) {
	newToken := uuid.New().String()

	oldToken, err := s.SessionRepository.FindExistTokenByUserUuid(ctx, user.Uuid)
	if err != nil {
		return "", "", err
	}

	// if token already exists, it is double login
	if oldToken != "" {
		// delete existing session in order to restrict old session
		if err := s.SessionRepository.DeleteSessionByUserUuid(ctx, user.Uuid); err != nil {
			return "", "", err
		}
	}

	// create new session
	if err = s.SessionRepository.CreateSession(ctx, &model.Session{
		Token:     newToken,
		UserUuid:  user.Uuid,
		CompanyId: user.CompanyId,
	}); err != nil {
		return "", "", err
	}

	return newToken, oldToken, nil
}
