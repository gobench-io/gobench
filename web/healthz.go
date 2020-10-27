package web

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) healthz(w http.ResponseWriter, r *http.Request) {
	err := h.s.PingDb()
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	// ping
	render.Respond(w, r, render.M{
		"hostname": h.s.GetHostname(),
	})
}
