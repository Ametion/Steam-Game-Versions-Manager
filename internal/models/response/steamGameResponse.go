package response

type SteamGameResponse struct {
	Data map[string]struct {
		Depots struct {
			Branches map[string]struct {
				BuildId string `json:"buildid"`
			} `json:"branches"`
		} `json:"depots"`
		Common struct {
			Name string `json:"name"`
		} `json:"common"`
	} `json:"data"`
	Status string `json:"status"`
}
