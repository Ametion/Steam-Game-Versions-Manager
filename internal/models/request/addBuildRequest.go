package request

type AddBuildBody struct {
	BuildId  string `json:"buildId"`
	IsTested bool   `json:"isTested"`
	GameId   uint   `json:"gameId"`
}
