package usersdto

import "housy/models"

type CreateUserRequest struct {
	Fullname string                      `json:"fullname" form:"fullname" validate:"required"`
	Username string                      `json:"username" form:"username" validate:"required"`
	Email    string                      `json:"email" form:"email" validate:"required"`
	Password string                      `json:"password" form:"password" validate:"required"`
	ListAsID int                         `json:"listasid" form:"listasid" validate:"required"`
	ListAs   models.ListAsIDUserResponse `json:"listas"`
	Gender   string                      `json:"gender" form:"gender" validate:"required"`
	Address  string                      `json:"address" form:"address" validate:"required"`
}

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	ListAsID int    `json:"listasid" form:"listasid"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
}
