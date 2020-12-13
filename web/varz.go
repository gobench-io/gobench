package web

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) varz(w http.ResponseWriter, r *http.Request) {
	varz, err := h.s.Varz()

	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	log.Printf("varz %+v\n", varz)

	vr := varzResponse{*varz}

	render.Respond(w, r, render.M{
		"hostname": h.s.GetHostname(),
	})

	err = render.Render(w, r, &vr)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
