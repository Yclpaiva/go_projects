package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Handler is a function that initializes all the routes for the API
func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware to check for authorization
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
