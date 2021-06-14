package user

import (
	"net/http"

	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

func Read(w http.ResponseWriter, r *http.Request) {
	user, _ := auth.User(r)
	// return model of authed user
	utils.JSONResponse(w, user, http.StatusOK)
	return
}
