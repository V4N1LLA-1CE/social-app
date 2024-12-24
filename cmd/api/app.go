package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.healthCheckHandler)

	return r
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
