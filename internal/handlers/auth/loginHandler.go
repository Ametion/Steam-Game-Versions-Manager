package authHandlers

import (
	"github.com/Ametion/gfx"
	"golang.org/x/crypto/bcrypt"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
	"steam-version-notificator/pkg/helpers/myJwt"
)

func LoginHandler(context *gfx.Context) {
	var body request.LoginBody
	var user databaseModels.User

	if bodyErr := context.SetBody(&body); bodyErr != nil {
		context.SendJSON(400, response.Response{
			Message: "something went wrong while set body from request",
			Code:    400,
		})
		return
	}

	userResult := database.GetDatabase().Where("login = ?", body.Login).First(&user)

	if userResult.Error != nil {
		context.SendJSON(400, response.Response{
			Message: "Something went wrong while search user with presented login",
			Code:    400,
		})
		return
	}

	hashError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if hashError != nil {
		context.SendJSON(400, response.Response{
			Message: "Something went wrong while check password",
			Code:    400,
		})
		return
	}

	token, tokenErr := myJwt.GenerateToken(&user)

	if tokenErr != nil {
		context.SendJSON(400, response.Response{
			Message: "Error while generate token for user",
			Code:    400,
		})
		return
	}

	context.SendJSON(200, response.LoginResponse{
		Token: token,
	})
}
