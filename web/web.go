package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/master"
	_ "github.com/gobench-io/gobench/web/statik"
	"github.com/rakyll/statik/fs"
)

type webKey string

type handler struct {
	logger logger.Logger
	s      *master.Master
	r      *chi.Mux
}

func (h *handler) db() *ent.Client {
	return h.s.DB()
}

// New return new router interface
func newHandler(s *master.Master, logger logger.Logger) *handler {
	h := &handler{
		s:      s,
		logger: logger,
	}

	// basic cors for more ideas, see:
	// https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// initialize the router
	r := chi.NewRouter()
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// rest for groups
	r.Route("/api/", func(r chi.Router) {
		r.Route("/groups", func(r chi.Router) {
			r.Get("/", h.listGroups) // GET /groups

			r.Route("/{groupID}", func(r chi.Router) {
				r.Use(h.groupCtx)
				r.Get("/", h.getGroup)
				r.Get("/graphs", h.getGroupGraphs)
			})
		})

		// rest for graphs
		r.Route("/graphs", func(r chi.Router) {
			r.Get("/", h.listGraphs) // GET /graphs

			r.Route("/{graphID}", func(r chi.Router) {
				r.Use(h.graphCtx)
				r.Get("/", h.getGraph)
				r.Get("/metrics", h.getGraphMetrics)
			})
		})

		// rest for metrics
		r.Route("/metrics", func(r chi.Router) {
			r.Get("/", h.listMetrics) // GET /metrics

			r.Route("/{metricID}", func(r chi.Router) {
				r.Use(h.metricCtx, h.timeCtx)

				r.Get("/", h.getMetric)
				r.Get("/counters", h.getMetricCounters)
				r.Get("/histograms", h.getMetricHistograms)
				r.Get("/gauges", h.getMetricGauges)
			})
		})

		// get the application
		r.Route("/applications", func(r chi.Router) {
			r.Get("/", h.listApplications)   // GET /applications
			r.Post("/", h.createApplication) // POST /applications

			r.Route("/{applicationID}", func(r chi.Router) {
				r.Use(h.applicationCtx)

				r.Get("/", h.getApplication)
				r.Get("/groups", h.getApplicationGroups)
				r.Put("/cancel", h.cancelApplication)
			})
		})
	})

	statikFS, err := fs.New()
	if err != nil {
		h.logger.Fatalw("failed read statikFS", "err", err)
	}

	r.Handle("/*", http.FileServer(statikFS))

	h.r = r

	return h
}

// Serve start a web server with given gobench server
func Serve(s *master.Master, logger logger.Logger) {
	h := newHandler(s, logger)

	portS := fmt.Sprintf(":%d", s.WebPort())

	logger.Infow("web server start", "port", portS)

	if err := http.ListenAndServe(portS, h.r); err != nil {
		logger.Fatalw("failed start HTTP server", "port", portS, "err", err)
	}
}
