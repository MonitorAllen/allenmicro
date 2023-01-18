package handler

import (
	"allenmicro/proto/user"
	"allenmicro/user-service-2/service"
	"context"
)

type handler struct {
}

func (h handler) UserRegistry(ctx context.Context, request *user.UserRegistryRequest, response *user.UserLoginResponse) error {
	*response = service.RegistryUser(request.GetUsername(), request.GetPassword(), request.GetEmail())
	return nil
}

func (h handler) UserLogin(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	//TODO implement me
	panic("implement me")
}

func Handler() user.UserHandler {
	return handler{}
}
