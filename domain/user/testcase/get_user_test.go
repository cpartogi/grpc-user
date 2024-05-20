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

func TestGetUser(t *testing.T) {
	mockRepo := new(mocks.UserRepoInterface)

	t.Run("Error repo get user by id", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, errors.New("failed")).Once()

		_, err := u.GetUser(context.Background(), &proto.GetUserRequest{
			UserId: "1",
		})

		assert.Error(t, err, "failed")
	})

	t.Run("Error data not found", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, nil).Once()

		_, err := u.GetUser(context.Background(), &proto.GetUserRequest{
			UserId: "1",
		})

		assert.Error(t, err, "rpc error: code = NotFound desc = not found")
	})

	t.Run("Success", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, &config.Config{
			Secret: config.SecretConfig{
				Key: "abc&1*~#^2^#s0^=A^^-test",
			},
		})

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "wD64IJxXVEyeZnZmS3Y=",
			PhoneNumber: "wD64IJxXVEyeZnZmS3Y=",
		}, nil).Once()

		_, err := u.GetUser(context.Background(), &proto.GetUserRequest{
			UserId: "a",
		})

		assert.NilError(t, err)
	})

	t.Run("Error metadata", func(t *testing.T) {
		u := usecase.NewUserUsecase(mockRepo, nil)
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("userid", "1"))

		mockRepo.On("GetUserById", mock.Anything, mock.Anything).Return(model.Users{
			Id:          "a",
			FullName:    "b",
			Email:       "c",
			PhoneNumber: "d",
		}, errors.New("failed")).Once()

		_, err := u.GetUser(ctx, &proto.GetUserRequest{
			UserId: "",
		})

		assert.Error(t, err, "rpc error: code = InvalidArgument desc = invalid metadata")
	})
}
