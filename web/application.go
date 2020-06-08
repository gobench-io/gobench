package web

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent/app"
)

func applicationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appID, err := strconv.Atoi(chi.URLParam(r, "applicationID"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		app, err := db.Application.
			Query().
			Where(app.ID(appID)).
			Only(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), webKey("application"), app)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func listApplications(w http.ResponseWriter, r *http.Request) {
	aps, err := db.Application.
		Query().
		All(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newApplicationListResponse(aps)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func createApplication(w http.ResponseWriter, r *http.Request) {
}

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

func getApplicationGroups(w http.ResponseWriter, r *http.Request) {
}
