package handler

import (
	"allenmicro/proto/user"
	"allenmicro/user-service/service"
	"context"
)

type handler struct {
}

func (h handler) UserRegistry(ctx context.Context, request *user.UserRegistryRequest, response *user.UserLoginResponse) error {
	*response = service.RegistryUser(request.GetUsername(), request.GetPassword(), request.GetEmail())
	return nil
}

func (h handler) UserLogin(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	*response = service.UserLogin(request.GetUsername(), request.GetPassword())
	return nil
}

func Handler() user.UserHandler {
	return handler{}
}
