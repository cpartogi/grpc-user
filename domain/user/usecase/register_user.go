package usecase

import (
	"context"
	"errors"
	"strings"
	"user-service/domain/user/model"
	"user-service/lib/pkg/utils"

	proto "user-service/pb/user"

	"user-service/lib/helper"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

func (u *UserUsecase) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (res *proto.RegisterUserResponse, err error) {

	invalidMessages, isValid := isDataValid(model.Users{
		FullName:     req.FullName,
		UserPassword: req.Password,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
	})

	if !isValid {
		errorMsg := strings.Join(invalidMessages, " , ")
		return nil, helper.Error(codes.InvalidArgument,
			"", errors.New(errorMsg))
	}

	//cek if email exist
	checkMail, err := u.userRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		if err != pg.ErrNoRows {
			return nil, helper.Error(codes.Internal,
				"", err)
		}
	}

	if checkMail.Email != "" {
		return nil, helper.Error(codes.AlreadyExists, "", err)
	}

	//insert to db
	id := uuid.New().String()

	pHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, helper.Error(codes.Internal, "", err)
	}

	req.Password = pHash

	userId, err := u.userRepo.InsertUser(ctx, model.Users{
		Id:           id,
		FullName:     req.FullName,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
		UserPassword: req.Password,
	})

	if err != nil {
		return nil, helper.Error(codes.Internal, "", err)
	}

	return &proto.RegisterUserResponse{
		Id: userId,
	}, nil
}
