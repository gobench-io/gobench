package web

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent"
)

func (h *handler) getEventLogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, ok := ctx.Value(webKey("application")).(*ent.Application)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	eventLogs, err := app.QueryEventLogs().All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newEventLogListResponse(eventLogs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
