package housedto

import "housy/models"

type CreateHouseRequest struct {
	Name      string                   `json:"name" form:"name" validate:"required"`
	CityID    int                      `json:"cityid" form:"cityid" validate:"required"`
	City      models.CityHouseResponse `json:"city"`
	Address   string                   `json:"address" form:"address" validate:"required"`
	Price     string                   `json:"price" form:"price" validate:"required"`
	TypeRent  string                   `json:"typerent" form:"typerent" validate:"required"`
	Ameneties string                   `json:"ameneties" form:"ameneties" validate:"required"`
	BedRoom   string                   `json:"bedroom" form:"bedroom" validate:"required"`
	BathRoom  string                   `json:"bathroom" form:"bathroom" validate:"required"`
}

// type UpdateUserRequest struct {
// 	Fullname string `json:"fullname" form:"fullname"`
// 	Username string `json:"username" form:"username"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// 	ListAsID int    `json:"listasid" form:"listasid"`
// 	Gender   string `json:"gender" form:"gender"`
// 	Address  string `json:"address" form:"address"`
// }
