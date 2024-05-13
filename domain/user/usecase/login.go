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

	"google.golang.org/grpc/codes"
)

func (u *UserUsecase) Login(ctx context.Context, req *proto.LoginRequest) (res *proto.LoginResponse, err error) {

	invalidMessages, isValid := isLoginValid(model.Users{
		UserPassword: req.Password,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return nil, helper.Error(codes.InvalidArgument, "", errors.New(errorMsg))
	}

	loginData, err := u.userRepo.GetUserByEmail(ctx, req.Email, "")

	if err != nil {
		return
	}

	if loginData.Id == "" {
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
			return nil, helper.Error(codes.Internal, "", err)
		}

		return nil, helper.Error(codes.InvalidArgument, "", errors.New(constant.PasswordWrong))
	}

	token, err := helper.GenerateTokenAndRefreshToken(loginData, &u.cfg.Token)
	if err != nil {
		return nil, helper.Error(codes.Internal, "", err)
	}

	err = u.userRepo.InsertUserLog(ctx, model.UserLogs{
		UserId:       loginData.Id,
		IsSuccess:    true,
		LoginMessage: "success",
	})

	if err != nil {
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
		return nil, helper.Error(codes.Internal, "", err)
	}

	return &proto.LoginResponse{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt.Format(time.RFC3339),
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiredAt: token.RefreshTokenExpiredAt.Format(time.RFC3339),
	}, nil
}
