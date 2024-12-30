package main

import (
	"github.com/V4N1LLA-1CE/social-app/internal/handlers/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() *chi.Mux {
	// initialise all handlers
	healthHandler := health.New(app.config.env)

	// init router
	r := chi.NewRouter()

	// global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", healthHandler.Check)
	})

	return r
}
