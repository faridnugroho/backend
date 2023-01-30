package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	// GetUser(ID int) (models.User, error)
	// CreateUser(user models.User) (models.User, error)
	// UpdateUser(user models.User) (models.User, error)
	// DeleteUser(user models.User) (models.User, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("House").Preload("User").Find(&transaction).Error

	return transaction, err
}
