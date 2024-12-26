package main

import (
	"fmt"
	"log"

	env "github.com/V4N1LLA-1CE/social-app/internal"
)

func main() {
	cfg := config{
		addr: fmt.Sprintf(":%s", env.GetString("PORT", "8080")),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
