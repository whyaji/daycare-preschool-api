package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/whyaji/daycare-preschool-api/config"
	"github.com/whyaji/daycare-preschool-api/internal/delivery/http"
	"github.com/whyaji/daycare-preschool-api/internal/repository"
	"github.com/whyaji/daycare-preschool-api/internal/usecase"
	"github.com/whyaji/daycare-preschool-api/pkg/database"
)

func main() {
	// Load env variables
	config.LoadEnv()
	cfg := config.GetConfig()

	// Connect to database
	db, err := database.ConnectDb(cfg)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Group routes
	api := app.Group("/api/v1")

	// User module
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(api, userUsecase)

	// Child module
	childRepo := repository.NewChildRepository(db)
	childUsecase := usecase.NewChildUsecase(childRepo)
	http.NewChildHandler(api, childUsecase)

	// Teacher Attendance module
	teacherAttendanceRepo := repository.NewTeacherAttendanceRepository(db)
	teacherAttendanceUsecase := usecase.NewTeacherAttendanceUsecase(teacherAttendanceRepo)
	http.NewTeacherAttendanceHandler(api, teacherAttendanceUsecase)

	// Start server
	log.Fatal(app.Listen(cfg.AppPort))
}
