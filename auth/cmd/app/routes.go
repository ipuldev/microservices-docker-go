package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/auth/authorize", app.authorize).Methods("GET")
	r.HandleFunc("/api/auth/token", app.token).Methods("GET")
	r.HandleFunc("/api/auth/signup", app.signup).Methods("POST")
	r.HandleFunc("/api/auth/login", app.login).Methods("POST")
	return r
}
