package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginBuyer(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "john"
	data["password"] = "123john123"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(j))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := make(map[string]string)
	err = json.Unmarshal([]byte(rr.Body.String()), &resp)
	if err != nil {
		t.Error(err.Error())
	}

	if _, ok := resp["auth_token"]; !ok {
		t.Error("auth token not present in response")
	}

	if val, ok := resp["auth_token"]; ok {
		if val == "" {
			t.Error("auth token is empty")
		}
	}

	authTokenBuyer = resp["auth_token"]
}

func TestLoginSeller(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "doe"
	data["password"] = "123doe123"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(j))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := make(map[string]string)
	err = json.Unmarshal([]byte(rr.Body.String()), &resp)
	if err != nil {
		t.Error(err.Error())
	}

	if _, ok := resp["auth_token"]; !ok {
		t.Error("auth token not present in response")
	}

	if val, ok := resp["auth_token"]; ok {
		if val == "" {
			t.Error("auth token is empty")
		}
	}

	authTokenSeller = resp["auth_token"]
}

func TestLoginInvalidCred(t *testing.T) {
	data := make(map[string]interface{})

	data["username"] = "doe"
	data["password"] = "123e123"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(j))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotAcceptable)
	}
}
