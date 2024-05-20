package test_user

import (
	"context"
	"errors"
	"testing"
	"user-service/config"
	"user-service/domain/user/mocks"
	"user-service/domain/user/model"
	"user-service/domain/user/usecase"
	"user-service/lib/helper"
	proto "user-service/pb/user"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestGetToken(t *testing.T) {
	mockRepo := new(mocks.UserRepoInterface)

	t.Run("Error forbidden", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.GetToken(context.Background(), &proto.GetTokenRequest{
			RefreshToken: "abc",
		})

		assert.Error(t, err, "rpc error: code = PermissionDenied desc = forbidden")
	})

	t.Run("Error repo get user by id", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Token: config.TokenConfig{
				Key:                "testprivatekey",
				Expiry:             10,
				RefreshTokenExpiry: 6,
			},
		})

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		}, &config.TokenConfig{
			Key:                "testprivatekey",
			Expiry:             10,
			RefreshTokenExpiry: 6,
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, errors.New("failed")).Once()

		_, err := u.GetToken(context.Background(), &proto.GetTokenRequest{
			RefreshToken: token.RefreshToken,
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error not found", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Token: config.TokenConfig{
				Key:                "testprivatekey",
				Expiry:             10,
				RefreshTokenExpiry: 6,
			},
		})

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		}, &config.TokenConfig{
			Key:                "testprivatekey",
			Expiry:             10,
			RefreshTokenExpiry: 6,
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "",
			FullName:    "",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.GetToken(context.Background(), &proto.GetTokenRequest{
			RefreshToken: token.RefreshToken,
		})

		assert.Error(t, err, "rpc error: code = NotFound desc = not found")
	})

	t.Run("Success generate token", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Token: config.TokenConfig{
				Key:                "testprivatekey",
				Expiry:             10,
				RefreshTokenExpiry: 6,
			},
		})

		token, _ := helper.GenerateTokenAndRefreshToken(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		}, &config.TokenConfig{
			Key:                "testprivatekey",
			Expiry:             10,
			RefreshTokenExpiry: 6,
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "123",
			FullName:    "a",
			Email:       "b",
			PhoneNumber: "c",
		}, nil).Once()

		_, err := u.GetToken(context.Background(), &proto.GetTokenRequest{
			RefreshToken: token.RefreshToken,
		})

		assert.NilError(t, err)
	})

}
