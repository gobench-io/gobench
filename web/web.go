package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/server"
	_ "github.com/gobench-io/gobench/web/statik"
	"github.com/rakyll/statik/fs"
)

type webKey string

var s *server.Server

func intServer(ws *server.Server) {
	s = ws
}

func db() *ent.Client {
	return s.DB()
}

// New return new router interface
func New(s *server.Server) *chi.Mux {
	intServer(s)

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
			r.Get("/", listGroups) // GET /groups

			r.Route("/{groupID}", func(r chi.Router) {
				r.Use(groupCtx)
				r.Get("/", getGroup)
				r.Get("/graphs", getGroupGraphs)
			})
		})

		// rest for graphs
		r.Route("/graphs", func(r chi.Router) {
			r.Get("/", listGraphs) // GET /graphs

			r.Route("/{graphID}", func(r chi.Router) {
				r.Use(graphCtx)
				r.Get("/", getGraph)
				r.Get("/metrics", getGraphMetrics)
			})
		})

		// rest for metrics
		r.Route("/metrics", func(r chi.Router) {
			r.Get("/", listMetrics) // GET /metrics

			r.Route("/{metricID}", func(r chi.Router) {
				r.Use(metricCtx, timeCtx)

				r.Get("/", getMetric)
				r.Get("/counters", getMetricCounters)
				r.Get("/histograms", getMetricHistograms)
				r.Get("/gauges", getMetricGauges)
			})
		})

		// get the application
		r.Route("/applications", func(r chi.Router) {
			r.Get("/", listApplications)   // GET /applications
			r.Post("/", createApplication) // POST /applications

			r.Route("/{applicationID}", func(r chi.Router) {
				r.Use(applicationCtx)

				r.Get("/", getApplication)
				r.Get("/groups", getApplicationGroups)
			})
		})
	})

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/*", http.FileServer(statikFS))

	return r
}

// Serve start a web server with given gobench server
func Serve(s *server.Server) {
	r := New(s)
	portS := fmt.Sprintf(":%d", s.WebPort())

	log.Printf("started the web server at port %s\n", portS)

	if err := http.ListenAndServe(portS, r); err != nil {
		log.Panicf("failed to start http server at port %s: %v", portS, err)
	}
}
