package web

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func getApplication(w http.ResponseWriter, r *http.Request) {
	apps, err := db.Application.Query().All(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if len(apps) == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	if len(apps) > 1 {
		log.Println("warning something went wrong, should not more thatn 2 applications", apps)
	}
	app := apps[0]

	render.DefaultResponder(w, r, app)
}
