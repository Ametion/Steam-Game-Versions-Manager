package databaseModels

type UserStatus string

const (
	Admin   UserStatus = "admin"
	Viewer  UserStatus = "viewer"
	Blocked UserStatus = "blocked"
)

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Login    string `gorm:"unique"`
	Password string `gorm:"longtext"`
	Status   UserStatus
}
