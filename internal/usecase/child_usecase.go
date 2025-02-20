package usecase

import (
	"encoding/json"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/repository"
)

type ChildUsecase interface {
	CreateChild(child *domain.Child) error
	CheckUserAdmin(userId uint) (bool, error)
	ValidateRequiredFields(requestData *domain.CreateChildRequest) []string
	ParseUserIds(userIds string) ([]domain.User, error)
	GetChild(id string) (*domain.Child, error)
}

type childUsecase struct {
	repo repository.ChildRepository
}

func NewChildUsecase(repo repository.ChildRepository) ChildUsecase {
	return &childUsecase{repo}
}

func (u *childUsecase) GetChild(id string) (*domain.Child, error) {
	return u.repo.GetChild(id)
}

func (u *childUsecase) CreateChild(child *domain.Child) error {
	return u.repo.Create(child)
}

func (u *childUsecase) CheckUserAdmin(userId uint) (bool, error) {
	user, err := u.repo.GetUserWithRoles(userId)
	if err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		if role.Name == "admin" {
			return true, nil
		}
	}
	return false, nil
}

func (u *childUsecase) ValidateRequiredFields(requestData *domain.CreateChildRequest) []string {
	var errors []string
	if requestData.Name == "" {
		errors = append(errors, "name is required")
	}
	if requestData.Nickname == "" {
		errors = append(errors, "nickname is required")
	}
	if requestData.BirthPlace == "" {
		errors = append(errors, "birthPlace is required")
	}
	if requestData.BirthDate == "" {
		errors = append(errors, "birthDate is required")
	}
	if requestData.Gender == "" {
		errors = append(errors, "gender is required")
	}
	if requestData.LivingWith == "" {
		errors = append(errors, "livingWith is required")
	}
	if requestData.RegisteredDate == "" {
		errors = append(errors, "registeredDate is required")
	}
	if requestData.Parents == "" || requestData.Parents == "[]" {
		errors = append(errors, "parents is required")
	}
	if requestData.Teachers == "" || requestData.Teachers == "[]" {
		errors = append(errors, "teachers is required")
	}
	return errors
}

func (u *childUsecase) ParseUserIds(userIds string) ([]domain.User, error) {
	var userIdsArray []uint
	if err := json.Unmarshal([]byte(userIds), &userIdsArray); err != nil {
		return nil, err
	}
	return u.repo.GetUsersByIds(userIdsArray)
}
