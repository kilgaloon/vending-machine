package auth

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/kilgaloon/atm/model/auth"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
)

// Middleware checks is user authed to api
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create the JWT key used to create the signature
		var jwtKey = []byte(os.Getenv("JWT_KEY"))

		// We can obtain the session token from the requests cookies, which come with every request
		c := r.Header.Get("X-Auth-Token")
		if c == "" {
			utils.JSONResponse(w, modelHttp.ErrorResponse{
				Message: "Invalid token",
			}, http.StatusUnauthorized)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c

		// Initialize a new instance of `Claims`
		claims := &auth.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if tkn != nil {
				if !tkn.Valid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
