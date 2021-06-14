package product

import (
	"net/http"

	"github.com/kilgaloon/atm/handlers/product/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
)

func Create(w http.ResponseWriter, r *http.Request) {
	req := proto.Create{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{}, http.StatusBadRequest)

		return
	}

	user, _ := auth.User(r)

	product := &model.Product{
		AmountAvailable: req.AmountAvailable,
		Cost:            req.Cost,
		ProductName:     req.ProductName,
		SellerID:        user.ID,
	}

	u, err := product.Create()
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
