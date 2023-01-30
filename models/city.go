package models

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

func (City) TableName() string {
	return "city"
}
