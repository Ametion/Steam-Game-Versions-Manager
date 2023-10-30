package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	databaseModels "steam-version-notificator/internal/database/models"
	"steam-version-notificator/internal/models/response"
	jwtModels "steam-version-notificator/pkg/helpers/myJwt/models"
	"strings"
)

func AuthorizationMiddleware() func(context *gin.Context) {

	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			context.JSON(403, response.Response{
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
			context.JSON(403, response.Response{
				Message: "Error while try to get jwt token",
				Code:    403,
			})
			context.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwtModels.Claims); ok && token.Valid {

			if claims.UserStatus == databaseModels.Blocked {
				context.JSON(403, response.Response{
					Message: "You need to ask for permission to use this service",
					Code:    403,
				})
				context.Abort()
				return
			}

			context.Set("user", claims.UserId)
			context.Set("userStatus", claims.UserStatus)
			context.Next()
		} else {
			context.JSON(403, response.Response{
				Message: "Error while check jwt token",
				Code:    403,
			})
			context.Abort()
			return
		}
	}
}
