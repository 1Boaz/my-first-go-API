package handelers

import (
	"github.com/1Boaz/my-first-go-API/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux)  {
	// global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/acount", func(router chi.Router) {
		// middleware for /acount route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}