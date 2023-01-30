package models

type ListAsID struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type ListAsIDUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (ListAsIDUserResponse) TableName() string {
	return "listasid"
}
