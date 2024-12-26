package main

import (
	"fmt"
	"log"

	"github.com/V4N1LLA-1CE/social-app/internal/env"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
)

func main() {
	cfg := config{
		addr: fmt.Sprintf(":%s", env.GetString("PORT", "8080")),
	}

	store := store.NewStore(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
