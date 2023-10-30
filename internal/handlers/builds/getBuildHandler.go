package buildHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetBuildHandler(context *gin.Context) {
	var dbBuild databaseModels.Build

	buildInfo := database.GetDatabase().Preload("Game").First(&dbBuild, context.Param("id"))

	if buildInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: buildInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	build := response.BuildResponse{
		Id:        dbBuild.Id,
		BuildId:   dbBuild.BuildId,
		InUse:     dbBuild.InUse,
		IsTested:  dbBuild.IsTested,
		GameId:    dbBuild.GameID,
		GameName:  dbBuild.Game.GameName,
		CreatedAt: dbBuild.CreatedAt,
		UpdatedAt: dbBuild.UpdatedAt,
	}

	context.JSON(200, build)
}
