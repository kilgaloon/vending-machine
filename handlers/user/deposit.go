package user

import (
	"net/http"

	"github.com/kilgaloon/atm/handlers/user/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

type DepositResponse struct {
	Amount uint64 `json:"amount"`
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	req := proto.Deposit{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{}, http.StatusBadRequest)

		return
	}

	user, _ := auth.User(r)
	amount, err := user.DepositAmount(model.Deposit(req.Amount))
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	resp := &DepositResponse{
		Amount: amount,
	}

	utils.JSONResponse(w, resp, http.StatusOK)
	return 
}
