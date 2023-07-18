package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/ent/graph"
)

func (h *handler) graphCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		graphID, err := strconv.Atoi(chi.URLParam(r, "graphID"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		g, err := h.db().Graph.
			Query().
			Where(graph.ID(graphID)).
			Only(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), webKey("graph"), g)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) listGraphs(w http.ResponseWriter, r *http.Request) {
	graphs, err := h.db().Graph.
		Query().
		All(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newGraphListResponse(graphs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getGraph(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	g, ok := ctx.Value(webKey("graph")).(*ent.Graph)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if err := render.Render(w, r, newGraphResponse(g)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getGraphMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	g, ok := ctx.Value(webKey("graph")).(*ent.Graph)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	ms, err := g.
		QueryMetrics().
		All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newMetricListResponse(ms)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
