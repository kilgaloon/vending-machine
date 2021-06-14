package product

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
	"github.com/kilgaloon/atm/handlers/user"
	"github.com/kilgaloon/atm/model"
	"github.com/kilgaloon/atm/utils"
	"github.com/kilgaloon/atm/utils/auth"

	userAuth "github.com/kilgaloon/atm/handlers/auth"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func cleanupUsers() {
	//setup for test
	db := utils.DBConnect()
	db.Exec("DELETE FROM products")
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

func getHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", userAuth.Auth).Methods("POST")
	r.HandleFunc("/products", List).Methods("GET")

	p := r.PathPrefix("/product").Subrouter()
	p.HandleFunc("", Create).Methods("POST")
	p.HandleFunc("/{id}", Update).Methods("PUT")
	p.HandleFunc("/{id}", Delete).Methods("DELETE")
	p.Use(auth.Middleware)
	p.Use(user.SellerMiddleware)

	buy := r.PathPrefix("/buy").Subrouter()
	buy.HandleFunc("", Buy).Methods("POST")
	buy.Use(user.BuyerMiddleware)

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

func getSellerId() uint {
	user := &model.User{
		Username: "test_seller",
	}

	user.Find()

	return user.ID
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
