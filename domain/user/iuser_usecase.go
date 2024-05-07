package user

import (
	"context"

	proto "user-service/pb/user"
)

type UserUsecaseInterface interface {
	RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (res *proto.RegiserUserResponse, err error)
}
