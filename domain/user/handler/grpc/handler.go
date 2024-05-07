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

func (h *Handler) RegisterUser(ctx context.Context, request *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {

	result, err := h.usecase.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *Handler) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {

	result, err := h.usecase.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
