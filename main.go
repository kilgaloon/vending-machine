package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kilgaloon/atm/handlers/user"
	"github.com/kilgaloon/atm/handlers/product"
	"github.com/kilgaloon/atm/model"
	"github.com/kilgaloon/atm/utils/auth"
	userAuth "github.com/kilgaloon/atm/handlers/auth"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// migrate models
	model.Migrate()

	r := mux.NewRouter()
	// user handlers
	r.HandleFunc("/user", user.Create).Methods("POST")
	r.HandleFunc("/user/login", userAuth.Auth).Methods("POST")

	u := r.PathPrefix("/user").Subrouter()
	u.HandleFunc("", user.Update).Methods("PUT")
	u.HandleFunc("", user.Read).Methods("GET")
	u.HandleFunc("", user.Delete).Methods("DELETE")
	u.Use(auth.Middleware)

	deposit := u.PathPrefix("/deposit").Subrouter()
	deposit.HandleFunc("", user.Deposit).Methods("PUT")
	deposit.Use(user.BuyerMiddleware)

	r.HandleFunc("/products", product.List).Methods("GET")

	p := r.PathPrefix("/product").Subrouter()
	p.HandleFunc("", product.Create).Methods("POST")
	p.HandleFunc("/{id}", product.Update).Methods("PUT")
	p.HandleFunc("/{id}", product.Delete).Methods("DELETE")
	p.Use(auth.Middleware)
	p.Use(user.SellerMiddleware)

	buy := r.PathPrefix("/buy").Subrouter()
	buy.HandleFunc("", product.Buy).Methods("POST")
	buy.Use(user.BuyerMiddleware)


	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
