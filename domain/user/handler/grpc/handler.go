package grpc_handler

import (
	"context"
	"user-service/domain/user"
	proto "user-service/pb/user"
)

type Handler struct {
	usecase user.UserUsecaseInterface
}

func NewHandler(usecase user.UserUsecaseInterface) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterUser(ctx context.Context, request *proto.RegisterUserRequest) (*proto.RegiserUserResponse, error) {

	result, err := h.usecase.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
