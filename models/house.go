package models

type House struct {
	ID        int    `json:"id"`
	Name      string `json:"name" form:"name" gorm:"type: varchar(255)"`
	CityID    int    `json:"-"`
	City      City   `json:"city"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Price     string `json:"price" gorm:"type: varchar(255)"`
	TypeRent  string `json:"typerent" gorm:"type: varchar(255)"`
	Ameneties string `json:"ameneties" gorm:"type: varchar(255)"`
	BedRoom   string `json:"bedroom" gorm:"type: varchar(255)"`
	BathRoom  string `json:"bathroom" gorm:"type: varchar(255)"`
}

func (House) TableName() string {
	return "houses"
}
