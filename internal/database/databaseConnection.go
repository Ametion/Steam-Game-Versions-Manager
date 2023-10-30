package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	databaseModels "steam-version-notificator/internal/database/models"
)

func ConnectDatabase() {
	db, connectionErr := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if connectionErr != nil {
		log.Fatalf("failed to connect database: %v", connectionErr.Error())
	}

	migrateErr := db.AutoMigrate(&databaseModels.Game{},
		&databaseModels.Build{}, &databaseModels.User{})
	if migrateErr != nil {
		log.Fatalf("failed to migrate: %v", migrateErr.Error())
	}
}

func GetDatabase() *gorm.DB {
	db, connectionErr := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if connectionErr != nil {
		log.Fatalf("failed to connect database: %v", connectionErr.Error())
	}

	return db
}
