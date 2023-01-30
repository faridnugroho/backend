package models

type ListAsRole struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

func (ListAsRole) TableName() string {
	return "listasroles"
}
