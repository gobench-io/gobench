package web

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) varz(w http.ResponseWriter, r *http.Request) {
	varz, err := h.s.Varz()

	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	vr := varzResponse{*varz}

	err = render.Render(w, r, &vr)
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
