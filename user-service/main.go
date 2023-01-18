package main

import (
	"allenmicro/proto/user"
	"allenmicro/user-service/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {

	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	service := micro.NewService(
		micro.Name("go.micro.allen.user"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	_ = user.RegisterUserHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
