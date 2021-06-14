package auth

import (
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/kilgaloon/atm/model"
	"github.com/kilgaloon/atm/model/auth"
)

// User Tries to resolve cookie and returns authed user
func User(r *http.Request) (*model.User, error) {
	// Create the JWT key used to create the signature
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	// Get the JWT string from the cookie
	tknStr := r.Header.Get("X-Auth-Token")
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
		return nil, errors.New("Invalid token")
	}

	if !tkn.Valid {
		return nil, errors.New("Invalid token")
	}

	user := &model.User{
		Username: claims.Username,
	}

	user, err = user.Find()
	if err != nil {
		return nil, errors.New("User with username doesn't exist")
	}

	return user, err
}
