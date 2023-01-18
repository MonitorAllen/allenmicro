package main

import (
	"allenmicro/proto/rand"
	"allenmicro/rand-service/handler"
	"allenmicro/rand-service/middleware"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {

	consulRegistry := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	middleware.InitSentinel()

	service := micro.NewService(
		micro.Name("go.micro.allen.rand"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(middleware.LimitingWrapper()),
	)

	service.Init()

	_ = rand.RegisterRandHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}
}
