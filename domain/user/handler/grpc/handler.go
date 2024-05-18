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
	functionName := "user-service.handler.RegisterUser"
	ctxHandler := helper.CheckRequestID(ctx)

	result, err := h.usecase.RegisterUser(ctxHandler, request)
	if err != nil {
		logger.Log(ctxHandler, functionName, err.Error(), request, result)
		return nil, err
	}

	logger.Log(ctxHandler, functionName, "", request, result)
	return result, nil
}

func (h *Handler) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	functionName := "user-service.handler.Login"
	result, err := h.usecase.Login(ctx, request)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), request, result)
		return nil, err
	}

	logger.Log(ctx, functionName, "", request, result)
	return result, nil
}

func (h *Handler) GetToken(ctx context.Context, request *proto.GetTokenRequest) (*proto.LoginResponse, error) {
	functionName := "user-service.handler.GetToken"
	result, err := h.usecase.GetToken(ctx, request)
	if err != nil {
		logger.Log(ctx, functionName, err.Error(), request, result)
		return nil, err
	}

	logger.Log(ctx, functionName, "", request, result)
	return result, nil
}
