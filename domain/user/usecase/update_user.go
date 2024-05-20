package usecase

import (
	"context"
	"errors"
	"strings"
	"user-service/domain/user/model"
	"user-service/lib/helper"
	"user-service/lib/pkg/encryption"
	logger "user-service/lib/pkg/logger"
	"user-service/lib/pkg/utils"
	proto "user-service/pb/user"

	"google.golang.org/grpc/codes"
)

func (u *UserUsecase) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (res *proto.UpdateUserResponse, err error) {
	functionName := "user-service.usecase.GetUser"

	//check metadata user
	userId, err := helper.CheckUserId(ctx)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), nil, nil)
		return res, helper.Error(codes.InvalidArgument, "", errors.New("invalid metadata"))
	}

	if req.UserId != userId {
		logger.Log(ctx, functionName, "permission denied", nil, nil)
		return res, helper.Error(codes.PermissionDenied, "", errors.New("permission denied"))
	}

	result, err := u.userRepo.GetUserById(ctx, userId)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}

	if result.Id == "" {
		logger.Log(ctx, functionName, "user not found", req, nil)
		return nil, helper.Error(codes.NotFound, "", errors.New("user not found"))
	}

	invalidMessages, isValid := isDataValidForUpdate(model.Users{
		FullName:     req.FullName,
		UserPassword: req.Password,
		PhoneNumber:  req.PhoneNumber,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		logger.Log(ctx, functionName, errorMsg, req, nil)
		return nil, helper.Error(codes.InvalidArgument, "", errors.New(errorMsg))
	}

	pHash, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}
	req.Password = pHash
	req.UserId = result.Id
	encryptPhone, _ := encryption.Encrypt(req.PhoneNumber, u.cfg.Secret)

	err = u.userRepo.UpdateUser(ctx, model.Users{
		Id:           result.Id,
		FullName:     req.FullName,
		PhoneNumber:  encryptPhone,
		UserPassword: req.Password,
	}, userId)

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, nil)
		return
	}

	logger.Log(ctx, functionName, "", req, nil)

	return &proto.UpdateUserResponse{
		Id: userId,
	}, nil

}
