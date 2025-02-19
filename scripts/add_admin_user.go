package scripts

import (
	"log"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunAddAdminUser(db *gorm.DB) {
	// Add admin user to the database
	log.Println("Adding admin user to the database")
	if adminUserError := addAdminUser(db); adminUserError != nil {
		log.Fatal("Adding admin user failed:", adminUserError)
	}
	log.Println("Admin user added successfully! 🚀")
}

// addAdminUser adds an admin user to the database
func addAdminUser(db *gorm.DB) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var adminRole domain.Role
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		return err
	}

	adminUser := domain.User{
		Name:     "Admin Puskaga",
		Email:    "adminpuskaga@uii.ac.id",
		Password: string(hashedPassword),
		Gender:   "male",
		Phone:    "-",
		Address:  "-",
		Roles:    []domain.Role{adminRole},
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	return nil
}
