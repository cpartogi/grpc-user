package user

import (
	"context"
	"user-service/domain/user/model"
)

type UserRepoInterface interface {
	InsertUser(ctx context.Context, req model.Users, requestId string) (userId string, err error)
	GetUserByEmail(ctx context.Context, email, requestId string) (res model.Users, err error)
	InsertUserLog(ctx context.Context, req model.UserLogs) (err error)
	UpsertUserToken(ctx context.Context, req model.UserToken) (err error)
}
