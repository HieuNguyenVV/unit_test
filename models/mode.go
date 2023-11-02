package models

type CreateUserRequest struct {
	ID       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}
