package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	token := authSeller()

	data := make(map[string]interface{})

	data["amount_available"] = 100
	data["cost"] = 50
	data["product_name"] = "gum"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
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

func TestCreateProductWithSameName(t *testing.T) {
	token := authSeller()

	data := make(map[string]interface{})

	data["amount_available"] = 100
	data["cost"] = 50
	data["product_name"] = "gum"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
	}

	req.Header.Set("X-Auth-Token", token)

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestBuyerProductCreate(t *testing.T) {
	token := authBuyer()

	data := make(map[string]interface{})

	data["amount_available"] = 100
	data["cost"] = 50
	data["product_name"] = "chocolate"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
	}

	req.Header.Set("X-Auth-Token", token)

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}