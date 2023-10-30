package gameHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
)

func AddGameHandler(context *gin.Context) {
	var body request.AddGameBody

	if bodyErr := context.ShouldBindJSON(&body); bodyErr != nil {
		context.JSON(400, response.Response{
			Message: bodyErr.Error(),
			Code:    400,
		})
		return
	}

	goReq := gorequest.New()

	var game response.SteamGameResponse

	goReq.Get("https://api.steamcmd.net/v1/info/" + body.GameId).EndStruct(&game)

	newGame := databaseModels.Game{
		GameId:    body.GameId,
		GameName:  game.Data[body.GameId].Common.Name,
		GameImage: body.GameImage,
		Builds: []databaseModels.Build{
			{
				BuildId: game.Data[body.GameId].Depots.Branches["public"].BuildId,
			},
		},
	}

	creationInfo := database.GetDatabase().Create(&newGame)

	if creationInfo.Error != nil {
		context.JSON(400, response.Response{
			Message: creationInfo.Error.Error(),
			Code:    400,
		})
		return
	}

	context.JSON(201, response.Response{
		Message: "Game added successfully",
		Code:    201,
	})
}
