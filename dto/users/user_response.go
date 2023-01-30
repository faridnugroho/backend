package usersdto

import "housy/models"

type UserResponse struct {
	ID       int                         `json:"id"`
	Fullname string                      `json:"fullname" form:"fullname" validate:"required"`
	Username string                      `json:"username" form:"username" validate:"required"`
	Email    string                      `json:"email" form:"email" validate:"required"`
	Password string                      `json:"password" form:"password" validate:"required"`
	ListAsID int                         `json:"-"`
	ListAs   models.ListAsIDUserResponse `json:"listas"`
	Gender   string                      `json:"gender" form:"gender" validate:"required"`
	Address  string                      `json:"address" form:"address" validate:"required"`
}
