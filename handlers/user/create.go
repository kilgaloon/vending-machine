package user

import (
	"net/http"

	"github.com/kilgaloon/atm/handlers/user/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	req := proto.Create{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{}, http.StatusBadRequest)

		return
	}

	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Role: model.Role(req.Role),
	}

	u, err := user.Create()
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
