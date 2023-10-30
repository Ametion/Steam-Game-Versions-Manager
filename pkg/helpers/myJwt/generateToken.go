package myJwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	databaseModels "steam-version-notificator/internal/database/models"
	jwtModels "steam-version-notificator/pkg/helpers/myJwt/models"
	"time"
)

func GenerateToken(usr *databaseModels.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &jwtModels.Claims{
		UserId:     usr.Id,
		UserStatus: usr.Status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	envErr := godotenv.Load()
	if envErr != nil {
		return "", envErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
