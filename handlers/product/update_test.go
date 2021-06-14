package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kilgaloon/atm/model"
)

func TestUpdate(t *testing.T) {
	token := authSeller()
	sellerID := getSellerId()

	p := &model.Product{
		SellerID: sellerID,
		ProductName: "gum",
	}

	prod, err := p.Find()
	if err != nil {
		t.Error(err.Error())
	}

	data := make(map[string]interface{})

	data["amount_available"] = 200
	data["cost"] = 10

	j, _ := json.Marshal(data)
	//

	endpoint := fmt.Sprintf("/product/%d", prod.ID)
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(j))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Auth-Token", token)

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
