package main

import (
	"net/http"

	"github.com/charmbracelet/log"

	"github.com/casali-dev/linkmap/internal/config"
	"github.com/casali-dev/linkmap/internal/db"
	"github.com/casali-dev/linkmap/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Unable to load env variables. Error: %s", err)
	}

	db := db.Init(cfg)

	s := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router.Handler(db),
	}

	log.Fatal(s.ListenAndServe())
}
