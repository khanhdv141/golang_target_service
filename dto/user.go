package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
