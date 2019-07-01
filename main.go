package main

import (
	"go-rest-api/controller"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
fmt.Println("Connected.")

	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
