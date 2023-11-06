package handlers

import (
	"fmt"
	"github.com/Ametion/gfx"
	"github.com/parnurzeal/gorequest"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
	"steam-version-notificator/pkg/helpers/discord"
)

func CheckVersionsHandler(context *gfx.Context) {
	goReq := gorequest.New()

	var dbGames []databaseModels.Game
	if err := database.GetDatabase().Preload("Builds").Find(&dbGames).Error; err != nil {
		context.SendJSON(400, response.Response{
			Message: err.Error(),
			Code:    400,
		})
		return
	}

	for i := range dbGames {
		var game response.SteamGameResponse
		_, _, errors := goReq.Get(fmt.Sprintf("https://api.steamcmd.net/v1/info/%s", dbGames[i].GameId)).EndStruct(&game)

		if errors != nil || len(errors) > 0 {
			context.SendJSON(400, response.Response{
				Message: errors[0].Error(),
				Code:    400,
			})
			return
		}

		latestBuildId := game.Data[dbGames[i].GameId].Depots.Branches["public"].BuildId
		buildExists := false

		for j := range dbGames[i].Builds {
			if dbGames[i].Builds[j].BuildId == latestBuildId {
				buildExists = true
				break
			}
		}

		if !buildExists {
			newBuild := databaseModels.Build{
				GameID:  dbGames[i].Id,
				BuildId: latestBuildId,
				InUse:   false,
			}

			if err := database.GetDatabase().Create(&newBuild).Error; err != nil {
				context.SendJSON(400, response.Response{
					Message: err.Error(),
					Code:    400,
				})
				return
			}

			notificationErr := discord.SendNotification("New version available for " + dbGames[i].GameName)
			if notificationErr != nil {
				context.SendJSON(400, response.Response{
					Message: notificationErr.Error(),
					Code:    400,
				})
				return
			}
		}
	}

	context.SendJSON(200, response.Response{
		Message: "All Games Checked",
		Code:    200,
	})
}
