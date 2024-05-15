package grpc_handler

import (
	"context"

	"user-service/domain/user"
	proto "user-service/pb/user"

	logger "user-service/lib/pkg/logger"
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
	functionName := "handler.RegisterUser"

	result, err := h.usecase.RegisterUser(ctx, request)
	if err != nil {
		logger.GetLogger(ctx, functionName, err.Error(), request, result)
		return nil, err
	}

	logger.GetLogger(ctx, functionName, "", request, result)
	return result, nil
}

func (h *Handler) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	functionName := "handler.Login"
	result, err := h.usecase.Login(ctx, request)
	if err != nil {
		logger.GetLogger(ctx, functionName, err.Error(), request, result)
		return nil, err
	}

	logger.GetLogger(ctx, functionName, "", request, result)
	return result, nil
}
