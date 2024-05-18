package usecase

import (
	"context"
	"errors"
	"strings"
	"time"
	"user-service/domain/user/model"
	"user-service/lib/constant"
	"user-service/lib/pkg/utils"
	proto "user-service/pb/user"

	"user-service/lib/helper"

	logger "user-service/lib/pkg/logger"

	"google.golang.org/grpc/codes"

	"user-service/lib/pkg/encryption"
)

func (u *UserUsecase) Login(ctx context.Context, req *proto.LoginRequest) (res *proto.LoginResponse, err error) {
	functionName := "usecase.RegisterUser"
	invalidMessages, isValid := isLoginValid(model.Users{
		UserPassword: req.Password,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		logger.Log(ctx, functionName, errorMsg, req, res)
		return nil, helper.Error(codes.InvalidArgument, "", errors.New(errorMsg))
	}

	encryptEmail, _ := encryption.Encrypt(req.Email, u.cfg.Secret)
	loginData, err := u.userRepo.GetUserByEmail(ctx, encryptEmail)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return
	}

	if loginData.Id == "" {
		logger.Log(ctx, functionName, "not found", req, res)
		return nil, helper.Error(codes.NotFound, "", err)
	}

	err = utils.CheckPasswordHash(req.Password, loginData.UserPassword)
	if err != nil {
		err = u.userRepo.InsertUserLog(ctx, model.UserLogs{
			UserId:       loginData.Id,
			IsSuccess:    false,
			LoginMessage: constant.PasswordWrong,
		})

		if err != nil {
			logger.Log(ctx, functionName, err.Error(), req, res)
			return nil, helper.Error(codes.Internal, "", err)
		}
		logger.Log(ctx, functionName, "invalid argument", req, res)
		return nil, helper.Error(codes.InvalidArgument, "", errors.New(constant.PasswordWrong))
	}

	token, err := helper.GenerateTokenAndRefreshToken(loginData, &u.cfg.Token)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return nil, helper.Error(codes.Internal, "", err)
	}

	err = u.userRepo.InsertUserLog(ctx, model.UserLogs{
		UserId:       loginData.Id,
		IsSuccess:    true,
		LoginMessage: "success",
	})

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return nil, helper.Error(codes.Internal, "", err)
	}

	err = u.userRepo.UpsertUserToken(ctx, model.UserToken{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt,
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiredAt: token.RefreshTokenExpiredAt,
	})

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return nil, helper.Error(codes.Internal, "", err)
	}

	logger.Log(ctx, functionName, "", req, nil)

	return &proto.LoginResponse{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt.Format(time.RFC3339),
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiredAt: token.RefreshTokenExpiredAt.Format(time.RFC3339),
	}, nil
}
