package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdate(t *testing.T) {
	data := make(map[string]interface{})

	data["password"] = "123john123new"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("PUT", "/user", bytes.NewBuffer(j))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Auth-Token", authTokenBuyer)

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestLoginWithOldCreds(t *testing.T) {
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

	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotAcceptable)
	}
}
