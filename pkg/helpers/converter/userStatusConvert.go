package converter

import (
	"errors"
	databaseModels "steam-version-notificator/internal/database/models"
)

func StringToUserStatus(s string) (databaseModels.UserStatus, error) {
	switch s {
	case string(databaseModels.Admin):
		return databaseModels.Admin, nil
	case string(databaseModels.Viewer):
		return databaseModels.Viewer, nil
	case string(databaseModels.Blocked):
		return databaseModels.Blocked, nil
	default:
		return "", errors.New("unknown user status")
	}
}
