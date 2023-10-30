package gameHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
)

func GetGamesHandler(context *gin.Context) {
	var games []response.GameResponse
	var dbGames []databaseModels.Game

	gamesInfo := database.GetDatabase().Preload("Builds").Find(&dbGames)

	if gamesInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: gamesInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	for i := range dbGames {
		game := response.GameResponse{
			GameId:        dbGames[i].Id,
			GameName:      dbGames[i].GameName,
			GameImage:     dbGames[i].GameImage,
			LatestBuildId: dbGames[i].Builds[len(dbGames[i].Builds)-1].BuildId,
			UpdatedAt:     dbGames[i].UpdatedAt,
		}

		games = append(games, game)
	}

	context.JSON(200, games)
}
