package migration

import (
	"log"

	"go-backend/config"

	"github.com/pressly/goose/v3"
)

func Migrate() {
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("Error setting goose dialect: %v", err)
	}

	if err := goose.Up(config.DB, "migration/scripts"); err != nil {
		log.Fatalf("Error running goose migrations: %v", err)
	}

	log.Println("Migration success ✅")
}