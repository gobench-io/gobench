package web

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/ent/counter"
	"github.com/gobench-io/gobench/v2/ent/gauge"
	"github.com/gobench-io/gobench/v2/ent/histogram"
	"github.com/gobench-io/gobench/v2/ent/metric"
)

// middleware to get metric with metricID in the url param
func (h *handler) metricCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// metricID is string
		metricID, err := strconv.Atoi(chi.URLParam(r, "metricID"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		metric, err := h.db().Metric.
			Query().
			Where(metric.ID(metricID)).
			Only(r.Context())

		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), webKey("metric"), metric)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, msInt*int64(time.Millisecond)), nil
}

// middleware to get `from` from query
func (h *handler) timeCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		fromS := r.URL.Query().Get("from")
		from, err := strconv.ParseInt(fromS, 10, 64)
		if err == nil {
			ctx = context.WithValue(r.Context(), webKey("from"), from)
		}

		endS := r.URL.Query().Get("end")
		end, err := strconv.ParseInt(endS, 10, 64)
		if err == nil {
			ctx = context.WithValue(ctx, webKey("end"), end)
		}

		if ctx == nil {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) listMetrics(w http.ResponseWriter, r *http.Request) {
	ms, err := h.db().Metric.
		Query().
		All(r.Context())
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}
	if err := render.RenderList(w, r, newMetricListResponse(ms)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getMetric(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(webKey("metric")).(*ent.Metric)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	if err := render.Render(w, r, newMetricResponse(m)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getMetricCounters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(webKey("metric")).(*ent.Metric)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	q := m.QueryCounters()

	from, ok := ctx.Value(webKey("from")).(int64)
	if ok {
		q.Where(counter.TimeGT(from))
	}
	end, ok := ctx.Value(webKey("end")).(int64)
	if ok {
		q.Where(counter.TimeLTE(end))
	}

	cs, err := q.All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newCounterListResponse(cs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getMetricHistograms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(webKey("metric")).(*ent.Metric)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	q := m.QueryHistograms()

	from, ok := ctx.Value(webKey("from")).(int64)
	if ok {
		q.Where(histogram.TimeGT(from))
	}
	end, ok := ctx.Value(webKey("end")).(int64)
	if ok {
		q.Where(histogram.TimeLTE(end))
	}

	hs, err := q.All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newHistogramListResponse(hs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (h *handler) getMetricGauges(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m, ok := ctx.Value(webKey("metric")).(*ent.Metric)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	q := m.QueryGauges()

	from, ok := ctx.Value(webKey("from")).(int64)
	if ok {
		q.Where(gauge.TimeGT(from))
	}
	end, ok := ctx.Value(webKey("end")).(int64)
	if ok {
		q.Where(gauge.TimeLTE(end))
	}

	gs, err := q.All(ctx)

	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	if err := render.RenderList(w, r, newGaugeListResponse(gs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
