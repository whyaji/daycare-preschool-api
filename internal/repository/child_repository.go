package repository

import (
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

type ChildRepository interface {
	Create(child *domain.Child) error
	GetUserWithRoles(userId uint) (domain.User, error)
	GetUsersByIds(userIds []uint) ([]domain.User, error)
	GetChild(id string) (*domain.Child, error)
}

type childRepository struct {
	db *gorm.DB
}

func NewChildRepository(db *gorm.DB) ChildRepository {
	return &childRepository{db}
}

func (r *childRepository) GetChild(id string) (*domain.Child, error) {
	var child domain.Child
	err := r.db.Preload("Teachers").Preload("Parents").Where("id = ?", id).First(&child).Error
	return &child, err
}

func (r *childRepository) Create(child *domain.Child) error {
	return r.db.Create(child).Error
}

func (r *childRepository) GetUserWithRoles(userId uint) (domain.User, error) {
	var user domain.User
	err := r.db.Preload("Roles").Where("id = ?", userId).First(&user).Error
	return user, err
}

func (r *childRepository) GetUsersByIds(userIds []uint) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Where("id IN ?", userIds).Find(&users).Error
	return users, err
}
