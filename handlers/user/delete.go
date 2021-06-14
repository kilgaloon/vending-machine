package user

import (
	"net/http"

	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	user, _ := auth.User(r)
	err := user.Delete()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	return
}
