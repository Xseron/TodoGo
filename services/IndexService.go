package services

import (
	"net/http"

	"github.com/Xseron/MyExtraGo/components"
	"github.com/Xseron/MyExtraGo/domain"
)

func IndexService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
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
