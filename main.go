package main

import (
	"go-rest-api/controller"
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	// Handle all preflight request
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")
	r.HandleFunc("/savings/add", controller.SavingHandler).
		Methods("POST")
	r.HandleFunc("/savings", controller.GetSavingsHandler).
		Methods("GET")

	fmt.Println("Server is ready and is listening at port :3000 . . .")
	log.Fatal(http.ListenAndServe(":3000", r))
}
