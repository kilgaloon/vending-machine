package user

import (
	"net/http"

	"github.com/kilgaloon/atm/handlers/user/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

func Update(w http.ResponseWriter, r *http.Request) {
	req := proto.Update{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{}, http.StatusBadRequest)

		return
	}

	authUser, _ := auth.User(r)

	user := &model.User{
		Username: authUser.Username,
	}

	user, err = user.Find()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	if (req.Password != "") {
		user.Password = req.Password
	}

	u, err := user.Update()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	// when user is created, return it to response
	utils.JSONResponse(w, u, http.StatusOK)

	return
}
