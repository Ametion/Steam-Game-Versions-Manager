package buildHandlers

import (
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetBuildsHandler(context *gfx.Context) {
	var builds []response.BuildResponse
	var dbBuilds []databaseModels.Build

	buildsInfo := database.GetDatabase().Preload("Game").Preload("LastModified").Find(&dbBuilds)

	if buildsInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: buildsInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	for i := range dbBuilds {
		build := response.BuildResponse{
			Id:           dbBuilds[i].Id,
			BuildId:      dbBuilds[i].BuildId,
			InUse:        dbBuilds[i].InUse,
			IsTested:     dbBuilds[i].IsTested,
			GameId:       dbBuilds[i].GameID,
			GameName:     dbBuilds[i].Game.GameName,
			LastModified: dbBuilds[i].LastModified.Login,
			CreatedAt:    dbBuilds[i].CreatedAt,
			UpdatedAt:    dbBuilds[i].UpdatedAt,
		}

		builds = append(builds, build)
	}

	context.SendJSON(200, builds)
}
