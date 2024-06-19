package handlers

import (
	"github.com/Xseron/TodoGo/services"
	"github.com/go-chi/chi/v5"
)

func mainHandlers(router *chi.Mux) {
	router.Get("/", services.IndexService)
}

func todoHandler(router *chi.Mux) {
	router.Put("/todo/add", services.AddTodoService)
	router.Delete("/todo/{id}", services.DeleteTodoElement)
}
