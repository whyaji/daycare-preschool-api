package scripts

import (
	"log"

	"github.com/whyaji/daycare-preschool-api/internal/domain"
	"gorm.io/gorm"
)

func RunAddWorkLocation(db *gorm.DB) {
	log.Println("Adding work location to the database")
	if workLocationError := addWorkLocation(db); workLocationError != nil {
		log.Fatal("Adding work location failed:", workLocationError)
	}
	log.Println("Work location added successfully! 🚀")
}

// addAdminUser adds an admin user to the database
func addWorkLocation(db *gorm.DB) error {
	workLocation := domain.WorkLocation{
		Name:      "Puskaga",
		Address:   "Jl. Kaliurang KM 14,5",
		Latitude:  -7.688025,
		Longitude: 110.414599,
	}

	if err := db.Create(&workLocation).Error; err != nil {
		return err
	}

	return nil
}
