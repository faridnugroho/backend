package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	CheckIn    time.Time `json:"checkin"`
	CheckOut   time.Time `json:"checkout"`
	HouseID    int       `json:"houseid"`
	House      House     `json:"house"`
	UserID     int       `json:"userid"`
	User       User      `json:"user"`
	Total      string    `json:"email" gorm:"type: varchar(255)"`
	Status     string    `json:"password" gorm:"type: varchar(255)"`
	Attachment string    `json:"gender" gorm:"type: varchar(255)"`
}
