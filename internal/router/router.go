package router

import (
	"net/http"

	"github.com/casali-dev/linkmap/internal/handlers"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{slug}", handlers.MakeRedirectHandler(db))

	return mux
}
