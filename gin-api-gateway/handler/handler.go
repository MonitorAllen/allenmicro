package handler

import (
	"allenmicro/proto/rand"
	"allenmicro/proto/user"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type APIHandler struct {
	randClient rand.RandService
	userClient user.UserService
}

func GetAPIHandler(randService rand.RandService, userService user.UserService) *APIHandler {
	return &APIHandler{
		randClient: randService,
		userClient: userService,
	}
}

func (s *APIHandler) Rand(ctx *gin.Context) {
	request := rand.RandRequest{}

	_ = ctx.ShouldBindQuery(&request)

	response, err := s.randClient.GetRand(context.Background(), &request)
	if err != nil {
		log.Println(err)
	}

	username, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"result":   response.GetResult(),
	})
}

func (s *APIHandler) RegistryUser(ctx *gin.Context) {
	request := user.UserRegistryRequest{}

	_ = ctx.ShouldBindJSON(&request)

	response, _ := s.userClient.UserRegistry(context.Background(), &request)

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.GetMessage(),
		"token":   response.GetToken(),
	})
}

func (s *APIHandler) UserLogin(ctx *gin.Context) {
	request := user.UserLoginRequest{}

	_ = ctx.ShouldBindJSON(&request)

	response, _ := s.userClient.UserLogin(context.Background(), &request)

	ctx.JSON(http.StatusOK, gin.H{
		"message": response.GetMessage(),
		"token":   response.GetToken(),
	})
}
