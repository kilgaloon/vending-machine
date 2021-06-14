package product

import (
	"net/http"

	"github.com/kilgaloon/atm/handlers/product/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

func Buy(w http.ResponseWriter, r *http.Request) {
	req := proto.Buy{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	user, _ := auth.User(r)

	buyRequest := model.BuyRequest{
		ProductID: req.ProductID,
		Amount:    req.Amount,
	}

	resp, err := user.Buy(buyRequest)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	// when user is created, return it to response
	utils.JSONResponse(w, resp, http.StatusOK)

	return
}
