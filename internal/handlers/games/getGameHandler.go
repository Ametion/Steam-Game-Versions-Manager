package gameHandlers

import (
	"github.com/Ametion/gfx"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetGameHandler(context *gfx.Context) {
	gameId := context.Param("id")

	var game databaseModels.Game

	gameInfo := database.GetDatabase().Preload("Builds").First(&game, gameId)

	if gameInfo.Error != nil {
		context.SendJSON(400, response.Response{
			Message: gameInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.SendJSON(200, response.GameResponse{
		GameId:        game.Id,
		GameName:      game.GameName,
		GameImage:     game.GameImage,
		LatestBuildId: game.Builds[len(game.Builds)-1].BuildId,
		UpdatedAt:     game.UpdatedAt,
	})
}
