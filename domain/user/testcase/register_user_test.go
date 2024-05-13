package test_user

import (
	"context"
	"errors"
	"testing"
	"user-service/domain/user/mocks"
	"user-service/domain/user/model"

	"user-service/domain/user/usecase"

	proto "user-service/pb/user"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestRegisterUser(t *testing.T) {
	mockRepo := new(mocks.UserRepoInterface)

	t.Run("Error invalid data", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
			Password:    "e",
		}, "a")

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = fullName must be at minimum 3 characters and maximum 60 characters , phoneNumber must be at minimum 10 characters and maximum 13 characters , phoneNumber must start with the Indonesia country code +62 , password must be minimum 6 characters and maximum 64 characters , invalid email address , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error email required", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "Full name and last name",
			PhoneNumber: "+628781122333",
			Password:    "AB781#kec",
		}, "a")

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = email is required , invalid email address")
	})

	t.Run("Error get email by user repo", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, errors.New("failed")).Once()

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "Full name and last name",
			PhoneNumber: "+628781122333",
			Password:    "AB781#kec",
			Email:       "abc@def.com",
		}, "a")

		assert.Error(t, err, "rpc error: code = Internal desc = failed")
	})

	t.Run("Error email already exist", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "c",
			PhoneNumber:  "d",
			UserPassword: "e",
		}, nil).Once()

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "Full name and last name",
			PhoneNumber: "+628781122333",
			Password:    "AB781#kec",
			Email:       "abc@def.com",
		}, "a")

		assert.Error(t, err, "rpc error: code = AlreadyExists desc = ")
	})

	t.Run("Error repo insert user", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{}, nil).Once()

		mockRepo.On("InsertUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("failed")).Once()

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "Full name and last name",
			PhoneNumber: "+628781122333",
			Password:    "AB781#kec",
			Email:       "abc@def.com",
		}, "a")

		assert.Error(t, err, "rpc error: code = Internal desc = failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{}, nil).Once()

		mockRepo.On("InsertUser", mock.Anything, mock.Anything, mock.Anything).Return("a", nil).Once()

		_, err := u.RegisterUser(context.Background(), &proto.RegisterUserRequest{
			FullName:    "Full name and last name",
			PhoneNumber: "+628781122333",
			Password:    "AB781#kec",
			Email:       "abc@def.com",
		}, "a")

		assert.NilError(t, err)
	})
}
