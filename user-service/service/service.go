package service

import (
	"allenmicro/common"
	"allenmicro/proto/user"
	"allenmicro/user-service/model"
)

func RegistryUser(username string, password string, email string) user.UserLoginResponse {
	userModel := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Channel:  1,
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

func UserLogin(username string, password string) user.UserLoginResponse {
	db := common.GetDB()

	userModel := &model.User{}

	err := db.Where("username = ? AND password = ?", username, password).Find(userModel).Error

	message := "Success"
	token := ""
	if err != nil {
		message = "Failed"
	} else {
		token = common.GetToken(userModel.Username, userModel.ID)
	}

	return user.UserLoginResponse{
		Message: message,
		Token:   token,
	}
}
