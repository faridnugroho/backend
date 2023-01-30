package models

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type CityHouseResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CityHouseResponse) TableName() string {
	return "city"
}
