package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/usecase"
	"github.com/whyaji/daycare-preschool-api/pkg/middleware"
	"github.com/whyaji/daycare-preschool-api/pkg/utils"
)

type ChildHandler struct {
	usecase usecase.ChildUsecase
}

func NewChildHandler(api fiber.Router, usecase usecase.ChildUsecase) *ChildHandler {
	handler := &ChildHandler{usecase}
	childGroup := api.Group("/childs")
	childGroup.Use(middleware.JWTProtected)
	childGroup.Post("/", handler.CreateChild)
	childGroup.Get("/:id", handler.GetChild)
	return handler
}

func (h *ChildHandler) GetChild(c *fiber.Ctx) error {
	id := c.Params("id")
	child, err := h.usecase.GetChild(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(child)
}

func (h *ChildHandler) CreateChild(c *fiber.Ctx) error {
	id := utils.GetUserIDFromJwt(c)
	isAdmin, err := h.usecase.CheckUserAdmin(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not allowed to create child"})
	}

	var errors []string

	// Create a map to first parse the raw JSON data
	var requestData domain.CreateChildRequest
	if err := c.BodyParser(&requestData); err != nil {
		errors = append(errors, err.Error())
	}

	// Validate required fields are not empty (except for AlergyInfo)
	errorsValidate := h.usecase.ValidateRequiredFields(&requestData)
	errors = append(errors, errorsValidate...)

	// Manually parse birthDate as time.Time
	birthDate, err := utils.ParseDateStringToTime(requestData.BirthDate)
	if err != nil {
		errors = append(errors, "birthDate "+err.Error())
	}

	// Manually parse registeredDate as time.Time
	registeredDate, err := utils.ParseDateStringToTime(requestData.RegisteredDate)
	if err != nil {
		errors = append(errors, "registeredDate "+err.Error())
	}

	// Parse the parents string into a slice of User structs
	parents, err := h.usecase.ParseUserIds(requestData.Parents)
	if err != nil {
		errors = append(errors, "parents "+err.Error())
	}

	// Parse the teachers string into a slice of User structs
	teachers, err := h.usecase.ParseUserIds(requestData.Teachers)
	if err != nil {
		errors = append(errors, "teachers "+err.Error())
	}

	// If there are any errors, return them
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}

	// Now parse the entire request into the Child struct
	var child domain.Child
	child.Name = requestData.Name
	child.Nickname = requestData.Nickname
	child.BirthPlace = requestData.BirthPlace
	child.BirthDate = *birthDate
	child.Gender = requestData.Gender
	child.AlergyInfo = requestData.AlergyInfo
	child.NumberOfSiblings = requestData.NumberOfSiblings
	child.LivingWith = requestData.LivingWith
	child.RegisteredDate = *registeredDate
	child.Notes = requestData.Notes
	child.Parents = parents
	child.Teachers = teachers

	if err := h.usecase.CreateChild(&child); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Child created"})
}
