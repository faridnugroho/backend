package housedto

import "housy/models"

type HouseResponse struct {
	ID        int                      `json:"id"`
	Name      string                   `json:"name" form:"name" validate:"required"`
	CityID    int                      `json:"cityid" form:"cityid" validate:"required"`
	City      models.CityHouseResponse `json:"-"`
	Address   string                   `json:"address" form:"address" validate:"required"`
	Price     string                   `json:"price" form:"price" validate:"required"`
	TypeRent  string                   `json:"typerent" form:"typerent" validate:"required"`
	Ameneties string                   `json:"ameneties" form:"ameneties" validate:"required"`
	BedRoom   string                   `json:"bedroom" form:"bedroom" validate:"required"`
	BathRoom  string                   `json:"bathroom" form:"bathroom" validate:"required"`
}

type GetHouseResponse struct {
	ID        int                      `json:"id"`
	Name      string                   `json:"name" form:"name" validate:"required"`
	CityID    int                      `json:"-"`
	City      models.CityHouseResponse `json:"city"`
	Address   string                   `json:"address" form:"address" validate:"required"`
	Price     string                   `json:"price" form:"price" validate:"required"`
	TypeRent  string                   `json:"typerent" form:"typerent" validate:"required"`
	Ameneties string                   `json:"ameneties" form:"ameneties" validate:"required"`
	BedRoom   string                   `json:"bedroom" form:"bedroom" validate:"required"`
	BathRoom  string                   `json:"bathroom" form:"bathroom" validate:"required"`
}
