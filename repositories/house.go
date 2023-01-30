package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type HouseRepository interface {
	FindHouse() ([]models.House, error)
	GetHouse(ID int) (models.House, error)
	CreateHouse(house models.House) (models.House, error)
	// UpdateUser(user models.User) (models.User, error)
	// DeleteUser(user models.User) (models.User, error)
}

func RepositoryHouse(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHouse() ([]models.House, error) {
	var house []models.House
	err := r.db.Preload("City").Find(&house).Error

	return house, err
}

func (r *repository) GetHouse(ID int) (models.House, error) {
	var house models.House
	err := r.db.Preload("City").First(&house, ID).Error

	return house, err
}

func (r *repository) CreateHouse(house models.House) (models.House, error) {
	err := r.db.Preload("City").Create(&house).Error

	return house, err
}

// func (r *repository) UpdateUser(user models.User) (models.User, error) {
// 	err := r.db.Save(&user).Error

// 	return user, err
// }

// func (r *repository) DeleteUser(user models.User) (models.User, error) {
// 	err := r.db.Delete(&user).Error

// 	return user, err
// }
