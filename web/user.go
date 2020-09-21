package web

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
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
