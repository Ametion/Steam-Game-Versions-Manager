package buildHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
)

func EditBuildHandler(context *gin.Context) {
	var body request.EditBuildBody
	var build databaseModels.Build
	buildId := context.Param("id")
	userId := context.GetUint("user")

	if bindErr := context.ShouldBindJSON(&body); bindErr != nil {
		context.JSON(400, response.Response{
			Message: bindErr.Error(),
			Code:    400,
		})
		return
	}

	buildInfo := database.GetDatabase().Where("id = ?", buildId).First(&build)

	if buildInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: buildInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	build.IsTested = body.IsTested
	build.LastModifiedId = userId

	updateInfo := database.GetDatabase().Save(&build)

	if updateInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: updateInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.JSON(200, response.Response{
		Message: "Build updated",
		Code:    200,
	})
}
