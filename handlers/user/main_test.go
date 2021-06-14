package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"

	userAuth "github.com/kilgaloon/atm/handlers/auth"
)

// store token from test and user it for other tests
var authTokenBuyer = ""
var authTokenSeller = ""

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func cleanupUsers() {
	//setup for test
	db := utils.DBConnect()
	db.Exec("DELETE FROM users")
}

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	cleanupUsers()

	createBuyer()
	createSeller()

	fmt.Print("Setup completed.\n")
}

func teardown() {

	fmt.Print("Teardown completed.\n")
}

func getHandler() http.Handler {

	r := mux.NewRouter()
	// user handlers
	r.HandleFunc("/user", Create).Methods("POST")
	r.HandleFunc("/user/login", userAuth.Auth).Methods("POST")

	u := r.PathPrefix("/user").Subrouter()
	u.HandleFunc("", Update).Methods("PUT")
	u.HandleFunc("", Read).Methods("GET")
	u.HandleFunc("", Delete).Methods("DELETE")
	u.Use(auth.Middleware)

	deposit := u.PathPrefix("/deposit").Subrouter()
	deposit.HandleFunc("", Deposit).Methods("PUT")
	deposit.Use(BuyerMiddleware)

	return r
}

func createBuyer() {
	data := make(map[string]interface{})

	data["username"] = "test_buyer"
	data["password"] = "buyer123"
	data["role"] = "buyer"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)
}

func createSeller() {
	data := make(map[string]interface{})

	data["username"] = "test_seller"
	data["password"] = "seller123"
	data["role"] = "seller"

	json, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)
}

func authBuyer() string {
	data := make(map[string]interface{})

	data["username"] = "test_buyer"
	data["password"] = "buyer123"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		log.Fatalf("Auth failed: handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := make(map[string]string)
	err = json.Unmarshal([]byte(rr.Body.String()), &resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := resp["auth_token"]; !ok {
		log.Fatalf("auth token not present in response")
	}

	if val, ok := resp["auth_token"]; ok {
		if val == "" {
			log.Fatalf("auth token is empty")
		}
	}

	return resp["auth_token"]
}

func authSeller() string {
	data := make(map[string]interface{})

	data["username"] = "test_seller"
	data["password"] = "seller123"

	j, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(getHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		log.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	resp := make(map[string]string)
	err = json.Unmarshal([]byte(rr.Body.String()), &resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := resp["auth_token"]; !ok {
		log.Fatalf("auth token not present in response")
	}

	if val, ok := resp["auth_token"]; ok {
		if val == "" {
			log.Fatalf("auth token is empty")
		}
	}

	return resp["auth_token"]
}
