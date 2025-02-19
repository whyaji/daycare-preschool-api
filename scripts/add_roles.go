package scripts

import (
	"log"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

func RunAddRoles(db *gorm.DB) {
	// Add roles to the database
	log.Println("Adding roles to the database")
	if roleError := addRoles(db); roleError != nil {
		log.Fatal("Adding roles failed:", roleError)
	}
	log.Println("Roles added successfully! 🚀")
}

// AddRoles adds roles to the database
func addRoles(db *gorm.DB) error {
	roles := []string{"teacher", "parent", "admin", "psychologist"}
	for _, role := range roles {
		db.Create(&domain.Role{Name: role})
	}
	return nil
}
