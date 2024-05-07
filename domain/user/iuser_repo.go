package user

import (
	"context"
	"user-service/domain/user/model"
)

type UserRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users) (userId string, err error)
	GetUserByEmail(ctx context.Context, email string) (res model.Users, err error)
}
