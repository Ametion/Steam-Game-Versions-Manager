package authHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetUsersHandler(context *gin.Context) {
	userStatus := context.MustGet("userStatus").(databaseModels.UserStatus)
	var users []response.UserResponse
	var dbUsers []databaseModels.User

	if userStatus != databaseModels.Admin {
		context.JSON(400, response.Response{
			Message: "You don't have permission to get users list",
			Code:    400,
		})
		return
	}

	usersInfo := database.GetDatabase().Find(&dbUsers)

	if usersInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: usersInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	for i := range dbUsers {
		user := response.UserResponse{
			Id:     dbUsers[i].Id,
			Login:  dbUsers[i].Login,
			Status: string(dbUsers[i].Status),
		}

		users = append(users, user)
	}

	context.JSON(200, users)
}
