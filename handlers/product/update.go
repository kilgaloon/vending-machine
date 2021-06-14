package product

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kilgaloon/atm/handlers/product/proto"
	"github.com/kilgaloon/atm/model"
	modelHttp "github.com/kilgaloon/atm/model/http"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"
	"gorm.io/gorm"
)

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := proto.Update{}
	err := req.FromJSON(r.Body)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{}, http.StatusBadRequest)

		return
	}

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	user, _ := auth.User(r)

	p := &model.Product{
		Model: gorm.Model{
			ID: uint(id),
		},
		SellerID: user.ID,
	}

	p, err = p.Find()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	p.AmountAvailable = req.AmountAvailable
	p.Cost = req.Cost

	p, err = p.Update()
	if err != nil {
		utils.JSONResponse(w, modelHttp.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	// when user is created, return it to response
	utils.JSONResponse(w, p, http.StatusOK)

	return
}
