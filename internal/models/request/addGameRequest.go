package request

type AddGameBody struct {
	GameId    string `json:"gameId"`
	GameImage string `json:"gameImage"`
}
