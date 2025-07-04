package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	URL       string    `gorm:"not null"                json:"url"`
	Slug      string    `gorm:"uniqueIndex;not null;size:20" json:"slug"`
	CreatedAt time.Time `gorm:"autoCreateTime"          json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"          json:"updatedAt"`
}

func (l *Link) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == "" {
		l.ID = uuid.New().String()
	}
	return
}

func CreateLink(url string, slug string) Link {
	return Link{
		URL:  url,
		Slug: slug,
	}
}
