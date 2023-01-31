package transactiondto

import (
	"housy/models"
	// "time"
)

type CreateTransactionRequest struct {
	// CheckIn    time.Time    `json:"checkin" form:"checkin" validate:"required"`
	// CheckOut   time.Time    `json:"checkout" form:"checkout" validate:"required"`
	HouseID    int          `json:"houseid" form:"houseid" validate:"required"`
	House      models.House `json:"house"`
	UserID     int          `json:"userid" form:"userid" validate:"required"`
	User       models.User  `json:"user"`
	Total      string       `json:"total" form:"total" validate:"required"`
	Status     string       `json:"status" form:"status" validate:"required"`
	Attachment string       `json:"attachment" form:"attachment" validate:"required"`
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
