package test_user

import (
	"context"
	"errors"
	"testing"
	"user-service/config"
	"user-service/domain/user/mocks"
	"user-service/domain/user/model"
	"user-service/domain/user/usecase"

	proto "user-service/pb/user"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
	"gotest.tools/assert"
)

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepoInterface)

	t.Run("Error invalid metadata", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		_, err := u.UpdateUser(context.Background(), &proto.UpdateUserRequest{
			UserId:      "1",
			FullName:    "2",
			Password:    "3",
			PhoneNumber: "4",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = invalid metadata")
	})

	t.Run("Error permission denied", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "1a"))

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "1",
			FullName:    "2",
			Password:    "3",
			PhoneNumber: "4",
		})

		assert.Error(t, err, "rpc error: code = PermissionDenied desc = permission denied")
	})

	t.Run("Error repo GetUserById", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "1a"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, errors.New("failed")).Once()

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "1a",
			FullName:    "2",
			Password:    "3",
			PhoneNumber: "4",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error data not found", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "1a"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "1a",
			FullName:    "2",
			Password:    "3",
			PhoneNumber: "4",
		})

		assert.Error(t, err, "rpc error: code = NotFound desc = user not found")
	})

	t.Run("Error invalid data for update", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "1a"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "1a",
			FullName:    "2",
			Password:    "3",
			PhoneNumber: "4",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = fullName must be at minimum 3 characters and maximum 60 characters , phoneNumber must be at minimum 10 characters and maximum 13 characters , phoneNumber must start with the Indonesia country code +62 , password must be minimum 6 characters and maximum 64 characters , password containing at least 1 capital characters AND 1 number AND 1 special (nonalpha-numeric) character")
	})

	t.Run("Error repo update user", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "a"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		mockRepo.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "a",
			FullName:    "Full Name ini",
			PhoneNumber: "+628781122333",
			Password:    "eASd@123",
		})
		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("user_id", "a"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		mockRepo.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		_, err := u.UpdateUser(ctx, &proto.UpdateUserRequest{
			UserId:      "a",
			FullName:    "Full Name ini",
			PhoneNumber: "+628781122333",
			Password:    "eASd@123",
		})

		assert.NilError(t, err)
	})
}
