package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"steam-version-notificator/internal/database"
	"steam-version-notificator/internal/handlers"
	authHandlers "steam-version-notificator/internal/handlers/auth"
	buildHandlers "steam-version-notificator/internal/handlers/builds"
	"steam-version-notificator/internal/handlers/games"
	"steam-version-notificator/pkg/helpers/middlewares"
)

func main() {
	database.ConnectDatabase()

	engine := gin.Default()

	myCors := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	engine.Use(myCors)

	auth := engine.Group("/auth")

	auth.POST("/login", authHandlers.LoginHandler)
	auth.POST("/register", authHandlers.RegisterHandler)

	api := engine.Group("/api")

	api.Use(middlewares.AuthorizationMiddleware())

	api.GET("/users", authHandlers.GetUsersHandler)
	//Route for change chosen user permission
	api.POST("/user", authHandlers.ChangeUserPermissionHandler)

	api.GET("/check", handlers.CheckVersionsHandler)

	//Games
	api.GET("/games", gameHandlers.GetGamesHandler)
	api.GET("/game/:id", gameHandlers.GetGameHandler)
	api.POST("/game", gameHandlers.AddGameHandler)
	api.DELETE("/game/:id", gameHandlers.DeleteGameHandler)

	//Builds
	api.GET("/builds", buildHandlers.GetBuildsHandler)
	api.GET("/build/:id", buildHandlers.GetBuildHandler)
	api.POST("/build", buildHandlers.AddBuildHandler)
	api.DELETE("/build/:id", buildHandlers.DeleteBuildHandler)
	api.PATCH("/build/:id", buildHandlers.EditBuildHandler)

	runErr := engine.Run(":7778")
	if runErr != nil {
		log.Fatalf("failed to run: %v", runErr.Error())
	}
}
