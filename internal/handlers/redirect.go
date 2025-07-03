package handlers

import (
	"net/http"

	"github.com/casali-dev/linkmap/internal/repositories"
	"gorm.io/gorm"
)

func MakeRedirectHandler(db *gorm.DB) http.HandlerFunc {
	repo := &repositories.GormLinkRepository{DB: db}

	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		link, err := repo.GetBySlug(slug)
		if err != nil {
			http.Error(w, "Link not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.URL, http.StatusFound)
	}
}
