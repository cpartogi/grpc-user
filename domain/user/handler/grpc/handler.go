package grpc_handler

import (
	"context"

	"user-service/domain/user"
	proto "user-service/pb/user"

	"user-service/lib/helper"
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
	requestId := helper.GenerateRandomString(16)

	result, err := h.usecase.RegisterUser(ctx, request, requestId)
	if err != nil {
		logger.GetLogger(functionName, err.Error(), requestId, request, result)
		return nil, err
	}

	logger.GetLogger(functionName, "", requestId, request, result)
	return result, nil
}

func (h *Handler) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	functionName := "handler.Login"
	requestId := helper.GenerateRandomString(16)
	result, err := h.usecase.Login(ctx, request)
	if err != nil {
		logger.GetLogger(functionName, err.Error(), requestId, request, result)
		return nil, err
	}

	logger.GetLogger(functionName, "", requestId, request, result)
	return result, nil
}
