package main

import (
	"log"

	"github.com/whyaji/daycare-preschool-api/config"
	"github.com/whyaji/daycare-preschool-api/pkg/database"
	"github.com/whyaji/daycare-preschool-api/scripts"
)

func main() {
	// Load env variables
	config.LoadEnv()
	cfg := config.GetConfig()

	// Connect to database
	db, err := database.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	scripts.RunAddRoles(db)
}
