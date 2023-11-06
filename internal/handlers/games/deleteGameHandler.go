package gameHandlers

import (
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func DeleteGameHandler(context *gfx.Context) {
	gameId := context.Param("id")

	var game databaseModels.Game

	gamesInfo := database.GetDatabase().Preload("Builds").Where("id = ?", gameId).
		First(&game)

	if gamesInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: gamesInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	for i := range game.Builds {
		deleteBuildInfo := database.GetDatabase().Delete(&game.Builds[i])

		if deleteBuildInfo.Error != nil {
			context.SendJSON(400, response.Response{
				Message: deleteBuildInfo.Error.Error(),
				Code:    400,
			})
			return
		}
	}

	deleteGameInfo := database.GetDatabase().Delete(&game)

	if deleteGameInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: deleteGameInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.SendJSON(200, response.Response{
		Message: "Game deleted successfully",
		Code:    200,
	})
}
