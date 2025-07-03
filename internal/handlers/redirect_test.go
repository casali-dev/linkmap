package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casali-dev/linkmap/internal/handlers"
	"github.com/casali-dev/linkmap/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	if err := db.AutoMigrate(&models.Link{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	link := models.Link{
		ID:   "test-id",
		URL:  "https://example.com",
		Slug: "abc123",
	}
	db.Create(&link)

	return db
}

func TestRedirectHandler(t *testing.T) {
	db := setupTestDB(t)
	handler := handlers.MakeRedirectHandler(db)

	req := httptest.NewRequest("GET", "/abc123", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("expected 302 Found, got %d", w.Code)
	}

	location := w.Header().Get("Location")
	if location != "https://example.com" {
		t.Errorf("expected redirect to https://example.com, got %s", location)
	}
}
