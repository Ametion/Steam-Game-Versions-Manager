package request

type ChangeUserPermissionBody struct {
	UserId uint   `json:"userId"`
	Status string `json:"status"`
}
