package authHandlers

import (
	"fmt"
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
	"steam-version-notificator/pkg/helpers/converter"
)

func ChangeUserPermissionHandler(context *gfx.Context) {
	var body request.ChangeUserPermissionBody
	userStatus := context.GetItem("userStatus").(databaseModels.UserStatus)

	if bodyErr := context.SetBody(&body); bodyErr != nil {
		context.SendJSON(400, "Something went wrong while set body from request")
		return
	}

	if userStatus == databaseModels.Viewer || userStatus == databaseModels.Blocked {
		context.SendJSON(400, response.Response{
			Message: "You don't have permission to change users permission",
			Code:    400,
		})
		return
	}

	var user databaseModels.User

	userInfo := database.GetDatabase().Where("id = ?", body.UserId).First(&user)

	if userInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: "Something went wrong while search user with presented id",
			Code:    400,
		})
		return
	}

	newStatus, statusErr := converter.StringToUserStatus(body.Status)

	if statusErr != nil {
		fmt.Println(statusErr)
		context.SendJSON(400, response.Response{
			Message: "Something went wrong while convert string to user status",
			Code:    400,
		})
		return
	}

	user.Status = newStatus

	updateInfo := database.GetDatabase().Save(&user)

	if updateInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: "Something went wrong while update user status",
			Code:    400,
		})
		return
	}

	context.SendJSON(200, response.Response{
		Message: "User status updated successfully",
		Code:    200,
	})
}
