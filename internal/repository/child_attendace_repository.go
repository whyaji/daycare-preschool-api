package repository

import (
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

type ChildAttendanceRepository interface {
	Create(childAttendance *domain.ChildAttendance) error
}

type childAttendanceRepository struct {
	db *gorm.DB
}

func NewChildAttendanceRepository(db *gorm.DB) ChildAttendanceRepository {
	return &childAttendanceRepository{db}
}

func (r *childAttendanceRepository) Create(childAttendance *domain.ChildAttendance) error {
	return r.db.Create(childAttendance).Error
}
