package buildHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func DeleteBuildHandler(context *gin.Context) {
	buildId := context.Param("id")

	deleteInfo := database.GetDatabase().Delete(&databaseModels.Build{}, buildId)

	if deleteInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: deleteInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.JSON(200, response.Response{
		Message: "Build deleted",
		Code:    200,
	})
}
