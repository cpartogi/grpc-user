package usecase

import (
	"context"
	"errors"
	"user-service/lib/helper"
	logger "user-service/lib/pkg/logger"
	proto "user-service/pb/user"

	"google.golang.org/grpc/codes"
)

func (u *UserUsecase) GetUser(ctx context.Context, req *proto.GetUserRequest) (res *proto.UserResponse, err error) {
	functionName := "user-service.usecase.GetUser"

	//check metadata user
	userId, err := helper.CheckUserId(ctx)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), nil, nil)
		return res, helper.Error(codes.InvalidArgument, "", errors.New("invalid metadata"))
	}

	result, err := u.userRepo.GetUserById(ctx, userId)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), nil, nil)
		return
	}

	if result.Id == "" {
		logger.Log(ctx, functionName, "not found", nil, nil)
		return res, helper.Error(codes.NotFound, "", errors.New("not found"))
	}

	res = &proto.UserResponse{
		Id:          result.Id,
		FullName:    result.FullName,
		PhoneNumber: result.PhoneNumber,
		Email:       result.Email,
	}

	return

}
