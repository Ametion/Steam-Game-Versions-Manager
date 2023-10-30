package databaseModels

import "time"

type Game struct {
	Id        uint    `gorm:"primaryKey"`
	GameId    string  `gorm:"unique"`
	GameName  string  `gorm:"unique"`
	GameImage string  `gorm:"longtext"`
	Builds    []Build `gorm:"foreignKey:GameID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
