package main

import (
	"allenmicro/gin-api-gateway/client"
	"allenmicro/gin-api-gateway/handler"
	"allenmicro/gin-api-gateway/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	// TODO
	randClient := client.GetRandClient()
	userClient := client.GetUserClient()
	apiHandler := handler.GetAPIHandler(randClient, userClient)

	user := v1.Group("/user")
	user.POST("/registry", apiHandler.RegistryUser)
	user.POST("/login", apiHandler.UserLogin)
	user.GET("/rand", middleware.AuthMiddleware(), apiHandler.Rand)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
