package db

import (
	"github.com/charmbracelet/log"

	"github.com/casali-dev/linkmap/internal/config"
	"github.com/casali-dev/linkmap/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Link{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	return db
}
