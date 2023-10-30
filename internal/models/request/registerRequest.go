package request

type RegisterBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
