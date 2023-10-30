package response

import "time"

type GameResponse struct {
	GameId        uint      `json:"gameId"`
	GameName      string    `json:"gameName"`
	GameImage     string    `json:"gameImage"`
	LatestBuildId string    `json:"latestBuildId"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
