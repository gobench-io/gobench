package web

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/ent/application"
	"github.com/gobench-io/gobench/v2/master"
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

func (h *handler) countApplications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	keyword := r.URL.Query().Get("keyword")
	q := h.db().Application.
		Query()
	if keyword != "" {
		q = q.Where(application.NameContains(keyword))
	}
	count, err := q.Count(ctx)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.Render(w, r, newCountApplicationResponse(count)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) listApplications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	keyword := r.URL.Query().Get("keyword")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	order := r.URL.Query().Get("order")
	isAsc, _ := strconv.ParseBool(r.URL.Query().Get("isAsc"))
	if limit == 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 || offset > 100 {
		offset = 0
	}
	if order == "" {
		order = "created_at"
	}
	q := h.db().Application.
		Query().
		WithTags()
	if keyword != "" {
		q = q.Where(application.NameContains(keyword))
	}
	if !isAsc {
		q = q.Order(
			ent.Desc(application.FieldCreatedAt),
		)
	} else {
		q = q.Order(
			ent.Asc(application.FieldCreatedAt),
		)
	}
	aps, err := q.Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newApplicationListResponse(aps)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func decode(e1, e2, e3 string) (string, string, string, error) {
	d1, err := base64.StdEncoding.DecodeString(e1)
	if err != nil {
		return "", "", "", err
	}
	d2, err := base64.StdEncoding.DecodeString(e2)
	if err != nil {
		return "", "", "", err
	}
	d3, err := base64.StdEncoding.DecodeString(e3)
	if err != nil {
		return "", "", "", err
	}
	return string(d1), string(d2), string(d3), nil
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

	scenario, gomod, gosum, err := decode(data.Scenario, data.Gomod, data.Gosum)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(errors.New("Invalid Data")))
		return
	}

	app, err := h.s.NewApplication(r.Context(), data.Name, scenario, gomod, gosum)

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

func (h *handler) getApplicationTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	ts, err := app.QueryTags().All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newTagListResponse(ts)); err != nil {
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
		if errors.Is(err, master.ErrAppIsFinished) {
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

func (h *handler) deleteApplication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if err := h.s.DeleteApplication(ctx, app.ID); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if err := render.Render(w, r, newApplicationResponse(nil)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) addApplicationTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	data := &tagRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	// get tags
	tag, err := h.s.GetTagByApplication(ctx, app, data.Name)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	// nochanges
	if tag != nil {
		if err := render.Render(w, r, newTagResponse(tag)); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
		return
	}

	tag, err = h.s.SetApplicationTag(ctx, app.ID, data.Name)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.Render(w, r, newTagResponse(tag)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) removeApplicationTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if app == nil {
		render.Render(w, r, ErrNotFoundRequest(errors.New("Application not found")))
		return
	}
	tagID, err := strconv.Atoi(chi.URLParam(r, "tagID"))
	if err != nil {
		render.Render(w, r, ErrNotFoundRequest(err))
		return
	}
	// get tags
	tag, err := h.s.GetTagByID(ctx, tagID)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if tag == nil {
		render.Render(w, r, ErrNotFoundRequest(errors.New("Tag not found")))
		return
	}
	err = h.s.RemoveApplicationTag(ctx, tag)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.Render(w, r, newTagResponse(tag)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getApplicationSystemLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	_, sl, _ := h.s.Logpaths(app.ID)

	http.ServeFile(w, r, sl)
}

func (h *handler) getApplicationUserLog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	_, _, ul := h.s.Logpaths(app.ID)

	http.ServeFile(w, r, ul)
}
