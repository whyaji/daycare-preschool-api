package repository

import (
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/pkg/types"
	"github.com/whyaji/daycare-preschool-api/pkg/utils"
	"gorm.io/gorm"
)

type TeacherAttendanceRepository interface {
	Create(teacherAttendance *domain.TeacherAttendance) error
	GetUserWithRoles(userId uint) (domain.User, error)
	GetLastTeacherAttendanceByUserId(userId uint) (domain.TeacherAttendance, error)
	UpdateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error
	GetAllWorkLocation() ([]domain.WorkLocation, error)
	GetTeacherAttendanceByUserId(userId uint, pagingationFilter types.PaginationFilter) ([]domain.TeacherAttendance, int, error)
}

type teacherAttendanceRepository struct {
	db *gorm.DB
}

func NewTeacherAttendanceRepository(db *gorm.DB) TeacherAttendanceRepository {
	return &teacherAttendanceRepository{db}
}

func (r *teacherAttendanceRepository) Create(teacherAttendance *domain.TeacherAttendance) error {
	return r.db.Create(teacherAttendance).Error
}

func (r *teacherAttendanceRepository) GetUserWithRoles(userId uint) (domain.User, error) {
	var user domain.User
	err := r.db.Preload("Roles").Where("id = ?", userId).First(&user).Error
	return user, err
}

func (r *teacherAttendanceRepository) GetLastTeacherAttendanceByUserId(userId uint) (domain.TeacherAttendance, error) {
	var teacherAttendance domain.TeacherAttendance
	err := r.db.Where("user_id = ?", userId).Order("created_at desc").First(&teacherAttendance).Error
	return teacherAttendance, err
}

// get pagination teacher attendance by user id
func (r *teacherAttendanceRepository) GetTeacherAttendanceByUserId(userId uint, paginationFilter types.PaginationFilter) ([]domain.TeacherAttendance, int, error) {
	var teacherAttendances []domain.TeacherAttendance
	var totalRecords int64

	// Start query with base condition
	query := r.db.Model(&domain.TeacherAttendance{}).Where("user_id = ?", userId)

	// Apply year and month filter based on column date
	query = utils.ApplyYearMonthFilter(query, paginationFilter.Filters, "date")

	// filter is pagination.Filters without year and month
	filters := paginationFilter.Filters
	delete(filters, "year")
	delete(filters, "month")

	// Apply dynamic filters
	query = utils.ApplyFilters(query, filters)

	// Get total records count
	err := query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Order(paginationFilter.OrderBy + " " + paginationFilter.Sort).
		Offset((paginationFilter.Page - 1) * paginationFilter.Limit).
		Limit(paginationFilter.Limit).Find(&teacherAttendances).Error

	if err != nil {
		return nil, 0, err
	}

	totalPages := int((totalRecords + int64(paginationFilter.Limit) - 1) / int64(paginationFilter.Limit))
	return teacherAttendances, totalPages, nil
}

func (r *teacherAttendanceRepository) UpdateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error {
	return r.db.Save(teacherAttendance).Error
}

// GetAllWorkLocation gets all work location
func (r *teacherAttendanceRepository) GetAllWorkLocation() ([]domain.WorkLocation, error) {
	var workLocations []domain.WorkLocation
	err := r.db.Find(&workLocations).Error
	return workLocations, err
}
