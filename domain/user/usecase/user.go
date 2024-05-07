package usecase

import "user-service/domain/user"

type UserUsecase struct {
	userRepo user.UserRepoInterface
}

func NewUserUsecase(userRepo user.UserRepoInterface) user.UserUsecaseInterface {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
