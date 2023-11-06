package buildHandlers

import (
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
)

func EditBuildHandler(context *gfx.Context) {
	var body request.EditBuildBody
	var build databaseModels.Build
	buildId := context.Param("id")
	userId := context.GetItem("user").(uint)

	if bindErr := context.SetBody(&body); bindErr != nil {
		context.SendJSON(400, response.Response{
			Message: bindErr.Error(),
			Code:    400,
		})
		return
	}

	buildInfo := database.GetDatabase().Where("id = ?", buildId).First(&build)

	if buildInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: buildInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	build.IsTested = body.IsTested
	build.LastModifiedId = userId

	updateInfo := database.GetDatabase().Save(&build)

	if updateInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: updateInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.SendJSON(200, response.Response{
		Message: "Build updated",
		Code:    200,
	})
}
