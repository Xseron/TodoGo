package handlers

import (
	"net/http"

	"github.com/Xseron/MyExtraGo/services"
	"github.com/go-chi/chi/v5"
)

func mainHandlers(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		services.IndexService(w, r)
	})
	router.Put("/todo/add", func(w http.ResponseWriter, r *http.Request) {
		services.AddTodoService(w, r)
	})
	router.Delete("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		services.DeleteTodoElement(w, r)
	})
}
