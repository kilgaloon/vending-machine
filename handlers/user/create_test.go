package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSeller(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "john"
	data["password"] = "123john123"
	data["role"] = "seller"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateBuyer(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "doe"
	data["password"] = "123doe123"
	data["role"] = "buyer"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateWithExistingUsername(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "john"
	data["password"] = "123john123"
	data["role"] = "seller"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
