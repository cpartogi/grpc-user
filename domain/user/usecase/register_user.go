package usecase

import (
	"context"
	"errors"
	"strings"
	"user-service/domain/user/model"
	"user-service/lib/pkg/utils"

	proto "user-service/pb/user"

	"user-service/lib/helper"

	logger "user-service/lib/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"

	"user-service/lib/pkg/encryption"
)

func (u *UserUsecase) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (res *proto.RegisterUserResponse, err error) {

	functionName := "user-service.usecase.RegisterUser"
	invalidMessages, isValid := isDataValid(model.Users{
		FullName:     req.FullName,
		UserPassword: req.Password,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		logger.Log(ctx, functionName, errorMsg, req, res)
		return nil, helper.Error(codes.InvalidArgument,
			"", errors.New(errorMsg))
	}

	//cek if email exist
	encryptEmail, _ := encryption.Encrypt(req.Email, u.cfg.Secret)
	checkMail, err := u.userRepo.GetUserByEmail(ctx, encryptEmail)

	if err != nil {
		if err != pg.ErrNoRows {
			logger.Log(ctx, functionName, err.Error(), req, res)
			return nil, helper.Error(codes.Internal,
				"", err)
		}
	}

	if checkMail.Email != "" {
		logger.Log(ctx, functionName, codes.AlreadyExists.String(), req, res)
		return nil, helper.Error(codes.AlreadyExists, "", err)
	}

	//insert to db
	id := uuid.New().String()

	pHash, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return nil, helper.Error(codes.Internal, "", err)
	}

	req.Password = pHash
	encryptPhone, _ := encryption.Encrypt(req.PhoneNumber, u.cfg.Secret)

	userId, err := u.userRepo.InsertUser(ctx, model.Users{
		Id:           id,
		FullName:     req.FullName,
		Email:        encryptEmail,
		PhoneNumber:  encryptPhone,
		UserPassword: req.Password,
	})

	if err != nil {
		logger.Log(ctx, functionName, err.Error(), req, res)
		return nil, helper.Error(codes.Internal, "", err)
	}

	logger.Log(ctx, functionName, "", req, res)

	return &proto.RegisterUserResponse{
		Id: userId,
	}, nil
}
