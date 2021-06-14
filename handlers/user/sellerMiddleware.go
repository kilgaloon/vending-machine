package user

import (
	"net/http"

	"github.com/kilgaloon/atm/utils/auth"
)

// Middleware checks is user authed to api
func SellerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, _ := auth.User(r)
		if !user.Role.IsSeller() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
