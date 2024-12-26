package main

import (
	"log"
	"net/http"
	"time"

	"github.com/V4N1LLA-1CE/social-app/internal/store"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
	store  store.Store
}

type config struct {
	addr string
}

func (app *application) serve(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server is listening on port %s\n", app.config.addr)

	return srv.ListenAndServe()
}
