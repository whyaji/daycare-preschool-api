package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"github.com/whyaji/daycare-preschool-api/internal/usecase"
	"github.com/whyaji/daycare-preschool-api/pkg/middleware"
	"github.com/whyaji/daycare-preschool-api/pkg/utils"
)

type TeacherAttendanceHandler struct {
	usecase usecase.TeacherAttendanceUsecase
}

func NewTeacherAttendanceHandler(api fiber.Router, usecase usecase.TeacherAttendanceUsecase) *TeacherAttendanceHandler {
	handler := &TeacherAttendanceHandler{usecase}
	teacherAttendanceGroup := api.Group("/teacher-attendances")
	teacherAttendanceGroup.Use(middleware.JWTProtected)
	teacherAttendanceGroup.Post("/clock-in", handler.ClockIn)
	teacherAttendanceGroup.Put("/clock-out", handler.ClockOut)
	teacherAttendanceGroup.Get("/last", handler.GetLastTeacherAttendance)
	return handler
}

func (h *TeacherAttendanceHandler) ClockIn(c *fiber.Ctx) error {
	id := utils.GetUserIDFromJwt(c)
	isTeacher, err := h.usecase.CheckUserTeacher(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isTeacher {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not allowed to create teacher attendance"})
	}

	lastTeacherAttendance, err := h.usecase.CheckLastIsClockedOut(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// timeNow := time.Now()
	// create example timeNow is 2025-02-20 07:45:00
	timeNow := time.Date(2025, 2, 20, 6, 40, 0, 0, time.Local)

	var errors []string

	// Create a map to first parse the raw JSON data
	var requestData domain.CreateTeacherAttendanceRequest
	if err := c.BodyParser(&requestData); err != nil {
		errors = append(errors, err.Error())
	}

	// Validate required fields are not empty
	errorsValidate := h.usecase.ValidateRequiredFieldsClock(&requestData)
	errors = append(errors, errorsValidate...)

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	isArroundWorkLocation, err := h.usecase.CheckIsInWorkLocation(requestData.Latitude, requestData.Longitude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isArroundWorkLocation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are not in work location"})
	}

	var morningOvertime int
	if requestData.IsOvertimeMorning {
		// 08:00 subtrack with timeNow in minutes
		// morningOvertime maximal is 60 minutes
		morningOvertime = 480 - (timeNow.Hour()*60 + timeNow.Minute())
		if morningOvertime > 60 {
			morningOvertime = 60
		}
	}

	// if lastTeacherAttendanceDate is today, then update the lastTeacherAttendance
	if lastTeacherAttendance != nil && lastTeacherAttendance.Date.Format("2006-01-02") == timeNow.Format("2006-01-02") && lastTeacherAttendance.ClockIn == nil {
		lastTeacherAttendance.ClockIn = &timeNow
		lastTeacherAttendance.OvertimeMorning = morningOvertime
		if err := h.usecase.UpdateTeacherAttendance(lastTeacherAttendance); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Teacher attendance updated successfully"})
	}

	teacherAttendance := &domain.TeacherAttendance{
		UserID:          uint(*id),
		Date:            timeNow,
		ClockIn:         &timeNow,
		OvertimeMorning: morningOvertime,
	}

	if err := h.usecase.CreateTeacherAttendance(teacherAttendance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Teacher attendance created successfully"})
}

func (h *TeacherAttendanceHandler) ClockOut(c *fiber.Ctx) error {
	id := utils.GetUserIDFromJwt(c)
	isTeacher, err := h.usecase.CheckUserTeacher(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isTeacher {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not allowed to create teacher attendance"})
	}

	teacherAttendance, err := h.usecase.CheckLastIsClockedIn(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// timeNow := time.Now()
	// create example timeNow is 2025-02-20 17:30:00
	timeNow := time.Date(2025, 2, 20, 16, 30, 0, 0, time.Local)

	var errors []string

	// Create a map to first parse the raw JSON data
	var requestData domain.CreateTeacherAttendanceRequest
	if err := c.BodyParser(&requestData); err != nil {
		errors = append(errors, err.Error())
	}

	// Validate required fields are not empty
	errorsValidate := h.usecase.ValidateRequiredFieldsClock(&requestData)
	errors = append(errors, errorsValidate...)

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors})
	}

	isArroundWorkLocation, err := h.usecase.CheckIsInWorkLocation(requestData.Latitude, requestData.Longitude)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isArroundWorkLocation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are not in work location"})
	}

	var afternoonOvertime int
	if requestData.IsOvertimeEvening {
		// 16:00 subtrack with timeNow in minutes
		// afternoonOvertime maximal is 60 minutes
		afternoonOvertime = (timeNow.Hour()*60 + timeNow.Minute()) - 960
		if afternoonOvertime > 60 {
			afternoonOvertime = 60
		}
	}

	var startHour float32
	startHourFlat := float32(teacherAttendance.ClockIn.Hour()) + float32(teacherAttendance.ClockIn.Minute())/60
	if teacherAttendance.ClockIn != nil {
		// if clockIn before 08:00, then startHour is start from 08:00
		if startHourFlat < 8 {
			startHour = 8
		} else {
			startHour = startHourFlat
		}
	}

	// if clockOut after 16:00 then endHour is 16:00
	var endHour float32
	endHourFlat := float32(timeNow.Hour()) + float32(timeNow.Minute())/60
	if endHourFlat > 16 {
		endHour = 16
	} else {
		endHour = endHourFlat
	}

	// calculate workHour
	workHour := endHour - startHour

	teacherAttendance.ClockOut = &timeNow
	teacherAttendance.OvertimeEvening = afternoonOvertime
	// workHour with 1 decimal
	teacherAttendance.WorkHour = float32(int(workHour*10)) / 10

	if err := h.usecase.UpdateTeacherAttendance(teacherAttendance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Teacher attendance updated successfully"})
}

func (h *TeacherAttendanceHandler) GetLastTeacherAttendance(c *fiber.Ctx) error {
	id := utils.GetUserIDFromJwt(c)
	isTeacher, err := h.usecase.CheckUserTeacher(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if !isTeacher {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not allowed to create teacher attendance"})
	}

	teacherAttendance, err := h.usecase.GetLastTeacherAttendanceByUserId(uint(*id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": teacherAttendance})
}
