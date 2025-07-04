package main

import (
	"net/http"

	"github.com/charmbracelet/log"

	"github.com/casali-dev/linkmap/internal/db"
	"github.com/casali-dev/linkmap/internal/router"
)

func main() {
	db := db.Init()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(db),
	}

	log.Fatal(s.ListenAndServe())
}
