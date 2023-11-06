package buildHandlers

import (
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetBuildHandler(context *gfx.Context) {
	var dbBuild databaseModels.Build

	buildInfo := database.GetDatabase().Preload("Game").First(&dbBuild, context.Param("id"))

	if buildInfo.Error != nil {
		context.SendJSON(400, response.Response{
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

	context.SendJSON(200, build)
}
