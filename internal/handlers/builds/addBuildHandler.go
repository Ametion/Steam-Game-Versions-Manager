package buildHandlers

import (
	"github.com/gin-gonic/gin"
	"steam-version-notificator/internal/database"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/request"
	"steam-version-notificator/internal/models/response"
)

func AddBuildHandler(context *gin.Context) {
	var body request.AddBuildBody
	userId := context.GetUint("user")

	if bindErr := context.ShouldBindJSON(&body); bindErr != nil {
		context.JSON(400, response.Response{
			Message: bindErr.Error(),
			Code:    400,
		})
		return
	}

	db := database.GetDatabase()

	// Start a new transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		context.JSON(500, response.Response{
			Message: tx.Error.Error(),
			Code:    500,
		})
		return
	}

	// Create new build
	newBuild := databaseModels.Build{
		BuildId:        body.BuildId,
		IsTested:       body.IsTested,
		InUse:          true,
		GameID:         body.GameId,
		LastModifiedId: userId,
	}

	if creationErr := tx.Create(&newBuild).Error; creationErr != nil {
		tx.Rollback()
		context.JSON(400, response.Response{
			Message: creationErr.Error(),
			Code:    400,
		})
		return
	}

	if updateErr := tx.Model(&databaseModels.Build{}).Where("game_id = ?", body.GameId).Not("id", newBuild.Id).Update("in_use", false).Error; updateErr != nil {
		tx.Rollback()
		context.JSON(400, response.Response{
			Message: updateErr.Error(),
			Code:    400,
		})
		return
	}

	if commitErr := tx.Commit().Error; commitErr != nil {
		context.JSON(500, response.Response{
			Message: commitErr.Error(),
			Code:    500,
		})
		return
	}

	context.JSON(201, response.Response{
		Message: "Build added successfully",
		Code:    201,
	})
}
