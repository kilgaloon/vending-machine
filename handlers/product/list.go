package product

import (
	"net/http"

	"github.com/kilgaloon/atm/model"
	"github.com/kilgaloon/atm/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
	var list []model.Product

	db := utils.DBConnect()
	db.Find(&list)

	utils.JSONResponse(w, list, http.StatusOK)
	return 
}
