package user

import (
	"context"
	"user-service/domain/user/model"
)

type UserRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users) (userId string, err error)
	GetUserByEmail(ctx context.Context, email string) (res model.Users, err error)
	InsertUserLog(ctx context.Context, req model.UserLogs) (err error)
	GetUserById(ctx context.Context, id string) (res model.Users, err error)
}
