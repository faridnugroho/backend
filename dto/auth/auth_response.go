package authdto

type LoginResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Username string `gorm:"type: varchar(255)" json:"username"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
