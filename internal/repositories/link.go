package repositories

import (
	"github.com/casali-dev/linkmap/internal/models"
	"gorm.io/gorm"
)

type LinkRepository interface {
	GetBySlug(slug string) (models.Link, error)
	Insert(link models.Link) error
}

type GormLinkRepository struct {
	DB *gorm.DB
}

func (r *GormLinkRepository) GetBySlug(slug string) (models.Link, error) {
	var link models.Link
	err := r.DB.First(&link, "slug = ?", slug).Error
	return link, err
}

func (r *GormLinkRepository) Insert(link models.Link) error {
	return r.DB.Create(&link).Error
}
