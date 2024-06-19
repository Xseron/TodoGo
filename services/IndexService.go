package services

import (
	"net/http"

	"github.com/Xseron/TodoGo/components"
	"github.com/Xseron/TodoGo/domain"
)

func IndexService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// TODO add chanel handling
	userData, ok := domain.UsersData[r.RemoteAddr]
	if !ok {
		userData = domain.UsersData.AddUser(r.RemoteAddr)
	}

	index := components.MainTemplate(userData)
	err := index.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
