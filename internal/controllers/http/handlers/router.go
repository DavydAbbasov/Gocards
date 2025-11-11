package handlers

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", pingHandler())

	return router
}
