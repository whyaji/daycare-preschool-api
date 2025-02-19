package repository

import (
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*domain.User, error)
	GetById(id uint) (*domain.User, error)
	Create(user *domain.User) error
	GetRegisteredByEmail(email string) (*domain.RegisteredEmail, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetById(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetRegisteredByEmail(email string) (*domain.RegisteredEmail, error) {
	var user domain.RegisteredEmail
	err := r.db.Preload("Roles").Where("email = ?", email).First(&user).Error
	return &user, err
}
