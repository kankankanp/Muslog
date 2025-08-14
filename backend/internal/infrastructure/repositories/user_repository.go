package repositories

import (
	"backend/internal/infrastructure/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id string) error {
	return r.DB.Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}