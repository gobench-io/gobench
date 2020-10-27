package web

import (
	"net/http"

	"github.com/go-chi/render"
)

// healthz checks status of the master by
// - ping the db
func (h *handler) healthz(w http.ResponseWriter, r *http.Request) {
	// ping
	err := h.s.PingDb()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.Respond(w, r, render.M{
		"hostname": h.s.GetHostname(),
	})
}
