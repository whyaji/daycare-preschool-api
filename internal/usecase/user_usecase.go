package usecase

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/whyaji/daycare-preschool-api/config"
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CheckRegisteredEmail(email string) (*domain.RegisteredEmail, error)
	Register(user *domain.User, registeredEmail *domain.RegisteredEmail) error
	AssignRegisteredAtEmail(registeredEmail domain.RegisteredEmail) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserById(id uint) (*domain.User, error)
	VerifyPassword(user *domain.User, password string) error
	Login(user *domain.User) (*string, *string, error)
	CheckUserAdmin(userId uint) (bool, error)
	RegisterEmail(registeredEmail *domain.RegisteredEmail) error
	ParseRoles(rolesString string) ([]domain.Role, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CheckRegisteredEmail(email string) (*domain.RegisteredEmail, error) {
	registered_email, err := u.repo.GetRegisteredByEmail(email)
	if err != nil {
		return nil, err
	}
	return registered_email, err
}

func (u *userUsecase) Register(user *domain.User, registeredEmail *domain.RegisteredEmail) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("could not hash password")
	}
	user.Password = string(hashedPassword)

	// Add roles from registered_email table to user.Roles
	user.Roles = append(user.Roles, registeredEmail.Roles...)
	return u.repo.Create(user)
}

func (u *userUsecase) AssignRegisteredAtEmail(registeredEmail domain.RegisteredEmail) error {
	now := time.Now()
	registeredEmail.RegisteredAt = &now
	return u.repo.UpdateRegisteredEmail(&registeredEmail)
}

func (u *userUsecase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserById(id uint) (*domain.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) VerifyPassword(user *domain.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(user *domain.User) (*string, *string, error) {
	// Generate JWT token
	exp := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return nil, nil, err
	}

	expString := exp.Format(time.RFC3339)

	return &t, &expString, nil
}

func (u *userUsecase) CheckUserAdmin(userId uint) (bool, error) {
	user, err := u.repo.GetByIdWithRoles(userId)
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

func (u *userUsecase) RegisterEmail(registeredEmail *domain.RegisteredEmail) error {
	return u.repo.CreateRegisteredEmail(registeredEmail)
}

func (u *userUsecase) ParseRoles(rolesString string) ([]domain.Role, error) {
	roles, err := u.repo.GetAllRoles()
	if err != nil {
		return nil, err
	}

	// rolesString "[1,2]" to []uint{1, 2}
	var roleIds []uint
	if err := json.Unmarshal([]byte(rolesString), &roleIds); err != nil {
		return nil, err
	}

	// Get roles from roleIds
	var returnRoles []domain.Role
	for _, roleId := range roleIds {
		for _, role := range roles {
			if role.ID == roleId {
				returnRoles = append(returnRoles, role)
			}
		}
	}

	return returnRoles, nil
}
