package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuyerDeposit(t *testing.T) {
	token := authBuyer()

	data := make(map[string]interface{})

	data["amount"] = 10

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("PUT", "/user/deposit", bytes.NewBuffer(j))
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

func TestDepositInvalidAmount(t *testing.T) {
	token := authBuyer()

	data := make(map[string]interface{})

	data["amount"] = 24

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("PUT", "/user/deposit", bytes.NewBuffer(j))
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

func TestSellerDeposit(t *testing.T) {
	token := authSeller()

	data := make(map[string]interface{})

	data["amount"] = 10

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("PUT", "/user/deposit", bytes.NewBuffer(j))
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
