package services

import (
	"github.com/casali-dev/linkmap/internal/models"
	"github.com/casali-dev/linkmap/internal/repositories"
	"github.com/casali-dev/linkmap/internal/utils"
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

func ShortenUrl(db *gorm.DB, url string) (string, error) {
	repo := &repositories.GormLinkRepository{DB: db}

	slug, err := utils.GenerateID(6)
	if err != nil {
		log.Fatalf("Error generating slug: %v", err)
		return "", err
	}

	link := models.CreateLink(url, slug)

	err = repo.Insert(link)
	if err != nil {
		return "", err
	}

	log.Info("Link created with success!")
	return slug, nil
}
