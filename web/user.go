package web

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

func (h *handler) userLogin(tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{"username": "admin"})
		if err != nil {
			render.Render(w, r, ErrInternalServer(err))
			return
		}
		err = render.Render(w, r, &accesstokenResponse{
			ID: tokenString,
		})
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}
}
