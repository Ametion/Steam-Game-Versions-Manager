package jwtModels

import (
	"github.com/dgrijalva/jwt-go"
	databaseModels "steam-version-notificator/internal/database/models"
)

type Claims struct {
	UserId     uint                      `json:"user_id"`
	UserStatus databaseModels.UserStatus `json:"user_status"`
	jwt.StandardClaims
}
