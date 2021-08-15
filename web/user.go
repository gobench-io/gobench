package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt"
)

const fifteenDays = 15 * 24 * time.Hour

func createToken(tokenAuth *jwtauth.JWTAuth) (token string, err error) {
	if tokenAuth == nil {
		token = "you-can-get-in"
		return
	}
	_, token, err = tokenAuth.Encode(jwt.MapClaims{
		"username": "admin",
		"exp":      jwtauth.ExpireIn(fifteenDays),
	})

	return
}

func (h *handler) userLogin(tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := new(accesstokenRequest)
		if err := render.Bind(r, d); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		if d.Username != h.adminUsername || d.Password != h.adminPassword {
			render.Render(w, r, ErrUnauthenticated(errors.New("invalid credentials")))
			return
		}

		tokenString, err := createToken(tokenAuth)

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
