package db

import (
	"log"

	"github.com/casali-dev/linkmap/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("linkmap.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Link{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	return db
}
