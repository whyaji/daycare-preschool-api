package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/usecase"
	"github.com/whyaji/daycare-preschool-api/pkg/middleware"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(api fiber.Router, usecase usecase.UserUsecase) {
	handler := &UserHandler{usecase}
	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Get("/", handler.Accessible)
	api.Get("/restricted", middleware.JWTProtected, handler.Restricted)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	registeredEmail, err := h.usecase.CheckRegisteredEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "email cannot be used to register"})
	}

	if _, err := h.usecase.GetUserByEmail(user.Email); err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "email is already in use"})
	}

	if err := h.usecase.Register(&user, registeredEmail); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := h.usecase.GetUserByEmail(input.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	if err := h.usecase.VerifyPassword(user, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	t, exp, err := h.usecase.Login(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token_data": fiber.Map{"token": t, "expired_at": exp}, "user_data": user})
}

func (h *UserHandler) Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func (h *UserHandler) Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	exp := claims["exp"].(float64)
	expiredAt := time.Unix(int64(exp), 0)

	id := claims["id"].(float64)
	userData, err := h.usecase.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user-not-found"})
	}
	return c.JSON(fiber.Map{"token_data": fiber.Map{"token": user.Raw, "expired_at": expiredAt}, "user_data": userData})
}
