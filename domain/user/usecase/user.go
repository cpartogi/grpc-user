package usecase

import (
	"user-service/config"
	"user-service/domain/user"
)

type UserUsecase struct {
	userRepo user.UserRepoInterface
	cfg      config.Config
}

func NewUserUsecase(userRepo user.UserRepoInterface, cfg config.Config) user.UserUsecaseInterface {
	return &UserUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
