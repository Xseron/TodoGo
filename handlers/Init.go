package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	addHandlers(router, mainHandlers)
	return router
}

func addHandlers(router *chi.Mux, handlersFunctions ...func(*chi.Mux)) {
	for _, handlerFunction := range handlersFunctions {
		handlerFunction(router)
	}
}
