package service

import (
	"allenmicro/common"
	"allenmicro/proto/user"
	"allenmicro/user-service-2/model"
)

func RegistryUser(username string, password string, email string) user.UserLoginResponse {
	userModel := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Channel:  2,
	}

	message := "Success"
	token := ""

	db := common.GetDB()

	if err := db.Create(&userModel).Error; err != nil {
		message = err.Error()
	} else {
		token = "Test Token"
	}

	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}
