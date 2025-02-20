package usecase

import (
	"fmt"
	"math"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/repository"
)

type TeacherAttendanceUsecase interface {
	CreateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error
	CheckUserTeacher(userId uint) (bool, error)
	ValidateRequiredFieldsClock(requestData *domain.CreateTeacherAttendanceRequest) []string
	CheckLastIsClockedOut(userId uint) (*domain.TeacherAttendance, error)
	CheckLastIsClockedIn(userId uint) (*domain.TeacherAttendance, error)
	UpdateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error
	GetLastTeacherAttendanceByUserId(userId uint) (domain.TeacherAttendance, error)
	CheckIsInWorkLocation(latitude, longitude float64) (bool, error)
}

type teacherAttendanceUsecase struct {
	repo repository.TeacherAttendanceRepository
}

func NewTeacherAttendanceUsecase(repo repository.TeacherAttendanceRepository) TeacherAttendanceUsecase {
	return &teacherAttendanceUsecase{repo}
}

func (u *teacherAttendanceUsecase) CreateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error {
	return u.repo.Create(teacherAttendance)
}

func (u *teacherAttendanceUsecase) CheckUserTeacher(userId uint) (bool, error) {
	user, err := u.repo.GetUserWithRoles(userId)
	if err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		if role.Name == "teacher" {
			return true, nil
		}
	}
	return false, nil
}

func (u *teacherAttendanceUsecase) ValidateRequiredFieldsClock(requestData *domain.CreateTeacherAttendanceRequest) []string {
	var errors []string
	if requestData.Latitude == 0 {
		errors = append(errors, "latitude is required")
	}
	if requestData.Longitude == 0 {
		errors = append(errors, "longitude is required")
	}
	return errors
}

func (u *teacherAttendanceUsecase) CheckLastIsClockedOut(userId uint) (*domain.TeacherAttendance, error) {
	teacherAttendance, err := u.repo.GetLastTeacherAttendanceByUserId(userId)
	if err != nil {
		return nil, nil
	}
	if teacherAttendance.ClockOut == nil && teacherAttendance.ClockIn != nil {
		return nil, fmt.Errorf("you have not clocked out yet")
	}
	return &teacherAttendance, nil
}

func (u *teacherAttendanceUsecase) CheckLastIsClockedIn(userId uint) (*domain.TeacherAttendance, error) {
	teacherAttendance, err := u.repo.GetLastTeacherAttendanceByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("you have not clocked in yet")
	}
	if teacherAttendance.ClockOut != nil && teacherAttendance.ClockIn == nil {
		return nil, fmt.Errorf("you have not clocked in yet")
	}
	return &teacherAttendance, nil
}

func (u *teacherAttendanceUsecase) UpdateTeacherAttendance(teacherAttendance *domain.TeacherAttendance) error {
	return u.repo.UpdateTeacherAttendance(teacherAttendance)
}

func (u *teacherAttendanceUsecase) GetLastTeacherAttendanceByUserId(userId uint) (domain.TeacherAttendance, error) {
	return u.repo.GetLastTeacherAttendanceByUserId(userId)
}

func (u *teacherAttendanceUsecase) CheckIsInWorkLocation(latitude, longitude float64) (bool, error) {
	const tolerance = 0.3 // 300 meters in kilometers

	workLocations, err := u.repo.GetAllWorkLocation()
	if err != nil {
		return false, err
	}
	for _, workLocation := range workLocations {
		if haversineDistance(workLocation.Latitude, workLocation.Longitude, latitude, longitude) <= tolerance {
			return true, nil
		}
	}
	return false, nil
}

func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Radius of Earth in kilometers
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)
	lat1Rad := lat1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
