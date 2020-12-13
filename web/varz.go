package web

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) varz(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, render.M{
		"hostname": h.s.GetHostname(),
	})
}
