package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Xseron/MyExtraGo/components"
	"github.com/Xseron/MyExtraGo/domain"
	"github.com/go-chi/chi/v5"
)

func AddTodoService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("name")
	statusStr := r.FormValue("status")

	var statusMap = map[string]domain.TodoStatus{
		"0": domain.Planing,
		"1": domain.Doing,
		"2": domain.Done,
		"3": domain.Droped,
	}

	status, ok := statusMap[statusStr]
	if !ok {
		http.Error(w, "Invalid status value", http.StatusBadRequest)
		return
	}

	userData := domain.UsersData.GetUser(r.RemoteAddr)
	element, err := userData.AddTodoElement(title, status)
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

	todoElement := components.OOBTodoElementTemplate(element)
	form := components.EmptyFormTemplate()

	err = todoElement.Render(r.Context(), w)
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
	_, err = userData.DeleteTodoElement(idParam)
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		errorComponent := components.GlobalErrorTemplate(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		errorComponent.Render(r.Context(), w)
	}
}
