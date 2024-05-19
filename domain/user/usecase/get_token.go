package usecase

import (
	"context"
	"errors"
	"time"
	"user-service/lib/helper"
	proto "user-service/pb/user"

	logger "user-service/lib/pkg/logger"

	"google.golang.org/grpc/codes"
)

func (u *UserUsecase) GetToken(ctx context.Context, req *proto.GetTokenRequest) (res *proto.LoginResponse, err error) {
	functionName := "user-service.usecase.GetToken"
	dataToken, expiredAt, err := helper.GetDataFromToken(req.RefreshToken, u.cfg)

	if err != nil {
		logger.Log(ctx, functionName, "forbidden", req, res)
		return res, helper.Error(codes.PermissionDenied, "", errors.New("forbidden"))
	}

	loginData, err := u.userRepo.GetUserById(ctx, dataToken.Id)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return
	}

	if loginData.Id == "" {
		logger.Log(ctx, functionName, "data not found", req, res)
		return res, helper.Error(codes.NotFound, "", errors.New("not found"))
	}

	token, err := helper.GenerateToken(loginData, &u.cfg.Token)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return
	}

	res = &proto.LoginResponse{
		Id:                    loginData.Id,
		Token:                 token.Token,
		TokenExpiredAt:        token.TokenExpiredAt.Format(time.RFC3339),
		RefreshToken:          req.RefreshToken,
		RefreshTokenExpiredAt: expiredAt,
	}

	logger.Log(ctx, functionName, "", nil, nil)

	return
}
