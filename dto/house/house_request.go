package housedto

import "housy/models"

type CreateHouseRequest struct {
	Name      string      `json:"name" form:"name" validate:"required"`
	CityID    int         `json:"cityid" form:"cityid" validate:"required"`
	City      models.City `json:"-"`
	Address   string      `json:"address" form:"address" validate:"required"`
	Price     string      `json:"price" form:"price" validate:"required"`
	TypeRent  string      `json:"typerent" form:"typerent" validate:"required"`
	Ameneties string      `json:"ameneties" form:"ameneties" validate:"required"`
	BedRoom   string      `json:"bedroom" form:"bedroom" validate:"required"`
	BathRoom  string      `json:"bathroom" form:"bathroom" validate:"required"`
}

type UpdateHouseRequest struct {
	Name      string `json:"name" form:"name"`
	CityID    int    `json:"cityid" form:"cityid"`
	Address   string `json:"address" form:"address"`
	Price     string `json:"price" form:"price"`
	TypeRent  string `json:"typerent" form:"typerent"`
	Ameneties string `json:"ameneties" form:"ameneties"`
	BedRoom   string `json:"bedroom" form:"bedroom"`
	BathRoom  string `json:"bathroom" form:"bathroom"`
}
