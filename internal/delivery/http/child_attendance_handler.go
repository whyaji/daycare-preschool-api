package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/whyaji/daycare-preschool-api/internal/usecase"
	"github.com/whyaji/daycare-preschool-api/pkg/middleware"
)

type ChildAttendanceHandler struct {
	usecase usecase.ChildAttendanceUsecase
}

func NewChildAttendanceHandler(api fiber.Router, usecase usecase.ChildAttendanceUsecase) *ChildAttendanceHandler {
	handler := &ChildAttendanceHandler{usecase}
	childAttendanceGroup := api.Group("/child-attendances")
	childAttendanceGroup.Use(middleware.JWTProtected)
	childAttendanceGroup.Post("/", handler.ChildArrival)
	return handler
}

func (h *ChildAttendanceHandler) ChildArrival(c *fiber.Ctx) error {
	childIdString := c.FormValue("childId")

	childId, err := strconv.Atoi(childIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid childId",
		})
	}

	date := c.FormValue("date")
	arrival := c.FormValue("arrival")

	return h.usecase.ChildArrival(uint(childId), date, arrival)
}
