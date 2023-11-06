package main

import (
	"github.com/Ametion/gfx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"steam-version-notificator/internal/database"
	"steam-version-notificator/internal/handlers"
	authHandlers "steam-version-notificator/internal/handlers/auth"
	buildHandlers "steam-version-notificator/internal/handlers/builds"
	gameHandlers "steam-version-notificator/internal/handlers/games"
	"steam-version-notificator/pkg/helpers/middlewares"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("failed to load env: %v", envErr.Error())
	}

	database.ConnectDatabase()

	engine := gfx.NewGFXEngine()

	cors := gfx.CorsConfig{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	engine.UseCors(cors)

	auth := engine.Group("/auth")

	auth.Post("/login", authHandlers.LoginHandler)
	auth.Post("/register", authHandlers.RegisterHandler)

	api := engine.Group("/api")

	api.UseMiddleware(middlewares.AuthorizationMiddleware)

	api.Get("/users", authHandlers.GetUsersHandler)
	//Route for change chosen user permission
	api.Post("/user", authHandlers.ChangeUserPermissionHandler)

	api.Get("/check", handlers.CheckVersionsHandler)

	//Games
	api.Get("/games", gameHandlers.GetGamesHandler)
	api.Get("/game/:id", gameHandlers.GetGameHandler)
	api.Post("/game", gameHandlers.AddGameHandler)
	api.Delete("/game/:id", gameHandlers.DeleteGameHandler)

	//Builds
	api.Get("/builds", buildHandlers.GetBuildsHandler)
	api.Get("/build/:id", buildHandlers.GetBuildHandler)
	api.Post("/build", buildHandlers.AddBuildHandler)
	api.Delete("/build/:id", buildHandlers.DeleteBuildHandler)
	api.Patch("/build/:id", buildHandlers.EditBuildHandler)

	runErr := engine.Run(":" + os.Getenv("PORT"))
	if runErr != nil {
		log.Fatalf("failed to run: %v", runErr.Error())
	}
}
