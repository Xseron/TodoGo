package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Xseron/MyExtraGo/components"
	"github.com/Xseron/MyExtraGo/domain"
	"github.com/go-chi/chi/v5"
)

func IndexService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var userData *domain.TodoData

	if data, ok := domain.UsersData[r.RemoteAddr]; ok {
		userData = data
	} else {
		userData = domain.UsersData.AddUser(r.RemoteAddr)
	}

	index := components.MainTemplate("Todo List", userData)
	err := index.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func AddTodoService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get form values
	title := r.FormValue("name")
	statusStr := r.FormValue("status")

	// Convert status string to TodoStatus type
	var status domain.TodoStatus
	switch statusStr {
	case "0":
		status = domain.Planing
	case "1":
		status = domain.Doing
	case "2":
		status = domain.Done
	case "3":
		status = domain.Droped
	default:
		http.Error(w, "Invalid status value", http.StatusBadRequest)
		return
	}

	userData := domain.UsersData.GetUser(r.RemoteAddr)
	element, err := userData.AddElement(title, status)
	if err != nil {
		formData := domain.NewFormData()
		formData.Errors["name"] = err.Error()
		formData.Values["name"] = title
		formData.Values["status"] = fmt.Sprintf("%d", status)

		form := components.ErrorTemplate(formData)
		http.Error(w, "", http.StatusUnprocessableEntity)
		form.Render(r.Context(), w)
		return
	}

	todoElements := components.OOBTodoElementTemplate(element)
	form := components.EmptyFormTemplate()

	err = todoElements.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}

	err = form.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func DeleteTodoElement(w http.ResponseWriter, r *http.Request) {
	idParam, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return
	}
	userData := domain.UsersData.GetUser(r.RemoteAddr)
	_, err = userData.DeleteElement(idParam)
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		errorComponent := components.GlobalErrorTemplate(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		errorComponent.Render(r.Context(), w)
	}
}
