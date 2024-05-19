package user

import (
	"context"

	proto "user-service/pb/user"
)

type UserUsecaseInterface interface {
	RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (res *proto.RegisterUserResponse, err error)
	Login(ctx context.Context, req *proto.LoginRequest) (res *proto.LoginResponse, err error)
	GetToken(ctx context.Context, req *proto.GetTokenRequest) (res *proto.LoginResponse, err error)
	GetUser(ctx context.Context, req *proto.GetUserRequest) (res *proto.UserResponse, err error)
}
