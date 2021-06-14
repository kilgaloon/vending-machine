package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kilgaloon/atm/handlers/auth/proto"
	"github.com/kilgaloon/atm/model"
	"github.com/kilgaloon/atm/model/auth"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	AuthToken string `json:"auth_token"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	req := proto.Login{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: "Invalid body",
		}, http.StatusBadRequest)

		return
	}

	u := &model.User{
		Username: req.Username,
	}

	u, err = u.Find()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: "Incorrect credentials",
		}, http.StatusNotAcceptable)

		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &auth.Claims{
		ID: u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	resp := &Response{
		AuthToken: tokenString,
	}

	utils.JSONResponse(w, resp, http.StatusOK)
	return
}
