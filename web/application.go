package web

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/server"
)

func (h *handler) applicationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appID, err := strconv.Atoi(chi.URLParam(r, "applicationID"))
		if err != nil {
			render.Render(w, r, ErrNotFoundRequest(err))
			return
		}

		app, err := h.db().Application.
			Query().
			Where(application.ID(appID)).
			Only(r.Context())

		if err != nil {
			render.Render(w, r, ErrNotFoundRequest(err))
			return
		}
		ctx := context.WithValue(r.Context(), webKey("application"), app)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) listApplications(w http.ResponseWriter, r *http.Request) {
	aps, err := h.db().Application.
		Query().
		Order(
			ent.Desc(application.FieldCreatedAt),
		).
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

func (h *handler) createApplication(w http.ResponseWriter, r *http.Request) {
	data := &applicationRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if data.Name == "" {
		render.Render(w, r, ErrInvalidRequest(errors.New("Name required")))
		return
	}
	if data.Scenario == "" {
		render.Render(w, r, ErrInvalidRequest(errors.New("Scenario required")))
		return
	}

	decScenario, err := base64.StdEncoding.DecodeString(data.Scenario)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(errors.New("Invalid Scenario")))
		return
	}
	scenario := string(decScenario)

	app, err := h.s.NewApplication(r.Context(), data.Name, scenario)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, newApplicationResponse(app))
}

func (h *handler) getApplication(w http.ResponseWriter, r *http.Request) {
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

func (h *handler) getApplicationGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	gs, err := app.QueryGroups().All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newGroupListResponse(gs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) cancelApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	na, err := h.s.CancelApplication(ctx, app.ID)

	// if err is ErrAppIsFinished, return 400 error
	// else return 500 error
	if err != nil {
		if errors.Is(err, server.ErrAppIsFinished) {
			render.Render(w, r, ErrAppIsFinished(err))
			return
		}
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.Render(w, r, newApplicationResponse(na)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
