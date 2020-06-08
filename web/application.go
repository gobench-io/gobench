package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
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
			Where(application.ID(appID)).
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
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if err := render.Render(w, r, newApplicationResponse(app)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func getApplicationGroups(w http.ResponseWriter, r *http.Request) {
}
