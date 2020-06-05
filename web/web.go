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

var db *ent.Client

func intDB(c *ent.Client) {
	db = c
}

type webKey string

// New return new router interface
func New(db *ent.Client) *chi.Mux {
	// save the db config
	intDB(db)

	// basic cors
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
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
	r.Use(middleware.Logger)
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
			r.Get("/", listGraphs) // GET /groups

			r.Route("/{graphID}", func(r chi.Router) {
				r.Use(graphCtx)
				r.Get("/", getGraph)
				r.Get("/metrics", getGraphMetrics)
			})
		})

		// rest for metrics
		r.Route("/metrics", func(r chi.Router) {
			r.Get("/", listMetrics) // GET /groups

			r.Route("/{metricID}", func(r chi.Router) {
				r.Use(metricCtx, timeCtx)

				r.Get("/", getMetric)
				r.Get("/counters", getMetricCounters)
				r.Get("/histograms", getMetricHistograms)
				r.Get("/gauges", getMetricGauges)
			})
		})

		// get the application
		r.Get("/application", getApplication)
	})

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/*", http.FileServer(statikFS))

	return r
}

// Serve start a web server at given port
// should be run in a go routine
func Serve(collect *server.Collect, port int) {
	r := New(collect.DB)

	portS := fmt.Sprintf(":%d", port)

	log.Printf("started the web server at port %s\n", portS)

	if err := http.ListenAndServe(portS, r); err != nil {
		log.Panicf("failed to start http server at port %s: %v", portS, err)
	}
}
