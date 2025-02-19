package scripts

import (
	"log"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// Run Migrations
	log.Println("Running Migrations")
	if migrationError := migrate(db); migrationError != nil {
		log.Fatal("Migration failed:", migrationError)
	}
	log.Println("Database migration completed successfully! 🚀")
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.RegisteredEmail{},
		&domain.User{},
		&domain.Role{},
		&domain.Child{},
		&domain.TeacherAttendance{},
		&domain.ChildAttendance{},
		&domain.ChildDiary{},
		&domain.ChildMeal{},
		&domain.ChildSleep{},
		&domain.ChildToilet{},
		&domain.ChildCondition{},
		&domain.LeaveRequest{},
	)
}
