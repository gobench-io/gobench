package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/gobench-io/gobench/v2/ent"
	"github.com/gobench-io/gobench/v2/logger"
	"github.com/gobench-io/gobench/v2/master"
)

const adminUsername = "admin"

type webKey string

type handler struct {
	logger        logger.Logger
	s             *master.Master
	adminUsername string
	adminPassword string
	r             *chi.Mux
}

//go:embed ui/gobench-ui/build
var staticFs embed.FS

func (h *handler) db() *ent.Client {
	return h.s.DbClient()
}

func setAuth(r chi.Router, tokenAuth *jwtauth.JWTAuth) {
	if tokenAuth != nil {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
	}
}

// New return new router interface
func newHandler(s *master.Master, adminPassword string, logger logger.Logger) *handler {
	h := &handler{
		s:             s,
		adminUsername: adminUsername,
		adminPassword: adminPassword,
		logger:        logger,
	}

	// basic cors for more ideas, see:
	// https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	var tokenAuth *jwtauth.JWTAuth
	if adminPassword != "" {
		tokenAuth = jwtauth.New("HS256", []byte(adminPassword), nil)
	}

	// initialize the router
	r := chi.NewRouter()
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthz", h.healthz)

	r.Get("/varz", h.varz)

	// rest for groups
	r.Route("/api/", func(r chi.Router) {
		r.Post("/users/login", h.userLogin(tokenAuth))

		r.Route("/groups", func(r chi.Router) {
			setAuth(r, tokenAuth)

			r.Get("/", h.listGroups) // GET /groups

			r.Route("/{groupID}", func(r chi.Router) {
				r.Use(h.groupCtx)
				r.Get("/", h.getGroup)
				r.Get("/graphs", h.getGroupGraphs)
			})
		})

		// rest for graphs
		r.Route("/graphs", func(r chi.Router) {
			setAuth(r, tokenAuth)

			r.Get("/", h.listGraphs) // GET /graphs

			r.Route("/{graphID}", func(r chi.Router) {
				r.Use(h.graphCtx)
				r.Get("/", h.getGraph)
				r.Get("/metrics", h.getGraphMetrics)
			})
		})

		// rest for metrics
		r.Route("/metrics", func(r chi.Router) {
			setAuth(r, tokenAuth)

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
			setAuth(r, tokenAuth)

			r.Get("/", h.listApplications)       // GET /applications
			r.Get("/count", h.countApplications) // GET /applications/count
			r.Post("/", h.createApplication)     // POST /applications

			r.Route("/{applicationID}", func(r chi.Router) {
				r.Use(h.applicationCtx)

				r.Get("/", h.getApplication)
				r.Delete("/", h.deleteApplication)
				r.Get("/groups", h.getApplicationGroups)
				r.Put("/cancel", h.cancelApplication)
				r.Route("/tags", func(r chi.Router) {
					r.Get("/", h.getApplicationTags)
					r.Put("/", h.addApplicationTag)
					r.Delete("/{tagID}", h.removeApplicationTag)
				})
				r.Get("/logs/system", h.getApplicationSystemLog)
				r.Get("/logs/user", h.getApplicationUserLog)
			})
		})
	})

	fsys, _ := fs.Sub(staticFs, "ui/gobench-ui/build")
	r.Handle("/*", http.FileServer(http.FS(fsys)))

	h.r = r

	return h
}

// Serve start a web server with given gobench server
func Serve(s *master.Master, adminPassword string, logger logger.Logger) {
	h := newHandler(s, adminPassword, logger)

	portS := fmt.Sprintf(":%d", s.WebPort())

	logger.Infow("web server start", "port", portS)

	if err := http.ListenAndServe(portS, h.r); err != nil {
		logger.Fatalw("failed start HTTP server", "port", portS, "err", err)
	}
}
