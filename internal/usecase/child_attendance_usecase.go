package usecase

import (
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/repository"
	"github.com/whyaji/daycare-preschool-api/pkg/utils"
)

type ChildAttendanceUsecase interface {
	ChildArrival(childId uint, date string, arrival string) error
}

type childAttendanceUsecase struct {
	repo repository.ChildAttendanceRepository
}

func NewChildAttendanceUsecase(repo repository.ChildAttendanceRepository) ChildAttendanceUsecase {
	return &childAttendanceUsecase{repo}
}

func (u *childAttendanceUsecase) ChildArrival(childId uint, date string, arrival string) error {
	parsedDate, err := utils.ParseDateStringToTime(date)
	if err != nil {
		return err
	}

	parsedArrival, err := utils.ParseDateTimeStringToTime(arrival)
	if err != nil {
		return err
	}

	childAttendance := domain.ChildAttendance{
		ChildID:         childId,
		Date:            *parsedDate,
		Arrival:         *parsedArrival,
		OvertimeMorning: utils.CalculateChildMorningOvertime(*parsedArrival),
	}
	return u.repo.Create(&childAttendance)
}
