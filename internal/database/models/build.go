package databaseModels

import "time"

type Build struct {
	Id             uint   `gorm:"primaryKey"`
	GameID         uint   `gorm:"not null"`
	Game           Game   `gorm:"foreignKey:GameID"`
	BuildId        string `gorm:"varchar(255)"`
	InUse          bool   `gorm:"not null;"`
	IsTested       bool   `gorm:"not null;"`
	LastModifiedId uint   `gorm:"not null;"`
	LastModified   User   `gorm:"foreignKey:LastModifiedId"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
