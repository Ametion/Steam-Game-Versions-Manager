package middlewares

import (
	"github.com/Ametion/gfx"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
	jwtModels "steam-version-notificator/pkg/helpers/myJwt/models"
	"strings"
)

func AuthorizationMiddleware(context *gfx.Context) {
	authHeader := context.Headers.Get("Authorization")

	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) != 2 {
		context.SendJSON(403, response.Response{
			Message: "Wrong bearerToken type, should be: Bearer TOKEN",
			Code:    403,
		})
		context.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(bearerToken[1], &jwtModels.Claims{}, func(token *jwt.Token) (interface{}, error) {
		envErr := godotenv.Load()
		if envErr != nil {
			return nil, envErr
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		context.SendJSON(403, response.Response{
			Message: "Error while try to get jwt token",
			Code:    403,
		})
		context.Abort()
		return
	}

	if claims, ok := token.Claims.(*jwtModels.Claims); ok && token.Valid {

		if claims.UserStatus == databaseModels.Blocked {
			context.SendJSON(403, response.Response{
				Message: "You need to ask for permission to use this service",
				Code:    403,
			})
			context.Abort()
			return
		}

		context.SetItem("user", claims.UserId)
		context.SetItem("userStatus", claims.UserStatus)
		context.Next()
	} else {
		context.SendJSON(403, response.Response{
			Message: "Error while check jwt token",
			Code:    403,
		})
		context.Abort()
		return
	}
}
