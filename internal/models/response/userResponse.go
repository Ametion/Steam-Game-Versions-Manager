package response

type UserResponse struct {
	Id     uint   `json:"id"`
	Login  string `json:"login"`
	Status string `json:"status"`
}
