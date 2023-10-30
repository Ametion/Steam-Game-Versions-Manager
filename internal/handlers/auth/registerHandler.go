package authHandlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
)

func RegisterHandler(context *gin.Context) {
	var body request.RegisterBody
	var user databaseModels.User

	if bodyErr := context.ShouldBindJSON(&body); bodyErr != nil {
		context.JSON(400, response.Response{
			Message: "Something went wrong while set body from request",
			Code:    400,
		})
		return
	}

	hashedPass, hashingErr := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	if hashingErr != nil {
		context.JSON(400, response.Response{
			Message: "Something went wrong while hashing password for user",
			Code:    400,
		})
		return
	}

	user.Login = body.Login
	user.Status = databaseModels.Blocked
	user.Password = string(hashedPass)

	userErr := database.GetDatabase().Create(&user)

	if userErr.Error != nil {
		context.JSON(400, response.Response{
			Message: "Something went wrong while register new user in database",
			Code:    400,
		})
		return
	}

	context.JSON(201, response.Response{
		Message: "User Registered Successfully",
		Code:    201,
	})
}
