package main

import (
	"fmt"
	"net/http"

	"github.com/Xseron/MyExtraGo/domain"
	"github.com/Xseron/MyExtraGo/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		fmt.Println(domain.UsersData)
		services.IndexService(w, r)
	})
	r.Put("/todo/add", func(w http.ResponseWriter, r *http.Request) {
		services.AddTodoService(w, r)
	})
	r.Delete("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		services.DeleteTodoElement(w, r)
	})
	http.ListenAndServe(":3000", r)
}
