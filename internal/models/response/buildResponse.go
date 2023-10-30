package response

import "time"

type BuildResponse struct {
	Id           uint      `json:"id"`
	BuildId      string    `json:"buildId"`
	InUse        bool      `json:"inUse"`
	IsTested     bool      `json:"isTested"`
	GameId       uint      `json:"gameId"`
	GameName     string    `json:"gameName"`
	LastModified string    `json:"lastModified"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
