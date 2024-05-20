package test_user

import (
	"context"
	"errors"
	"testing"
	"user-service/config"
	"user-service/domain/user/mocks"
	"user-service/domain/user/model"
	"user-service/domain/user/usecase"
	"user-service/lib/pkg/utils"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"

	proto "user-service/pb/user"
)

func TestLogin(t *testing.T) {
	mockRepo := new(mocks.UserRepoInterface)

	t.Run("Error invalid data", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "c",
			Password: "e",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = invalid email address , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error email required", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Password: "eASd@123",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = email is required , invalid email address")
	})

	t.Run("Error repo get user by email", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, errors.New("failed")).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: "eASd@123",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error data not found", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: "eASd@123",
		})

		assert.Error(t, err, "rpc error: code = NotFound desc = ")
	})

	t.Run("Error repo insert user log", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: "eASd@123",
		})

		assert.Error(t, err, "rpc error: code = Internal desc = failed")
	})

	t.Run("Error wrong password", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: "eASd@123",
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: "eASd@123",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = wrong password")
	})

	t.Run("Error repo insert user log password match", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Token: config.TokenConfig{
				Key:                "abce",
				Expiry:             10,
				RefreshTokenExpiry: 20,
			},
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		userPassword := "eASd@123"
		userPassHash, _ := utils.HashPassword(userPassword)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: userPassHash,
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: userPassword,
		})

		assert.Error(t, err, "rpc error: code = Internal desc = failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Token: config.TokenConfig{
				Key:                "abce",
				Expiry:             10,
				RefreshTokenExpiry: 20,
			},
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		userPassword := "eASd@123"
		userPassHash, _ := utils.HashPassword(userPassword)

		mockRepo.On("GetUserByEmail", mock.Anything, mock.Anything, mock.Anything).Return(model.Users{
			Id:           "a",
			FullName:     "b",
			Email:        "abc@def.com",
			PhoneNumber:  "d",
			UserPassword: userPassHash,
		}, nil).Once()

		mockRepo.On("InsertUserLog", mock.Anything, mock.Anything).Return(nil).Once()

		mockRepo.On("UpsertUserToken", mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.Login(context.Background(), &proto.LoginRequest{
			Email:    "abc@def.com",
			Password: userPassword,
		})

		assert.NilError(t, err)
	})
}
