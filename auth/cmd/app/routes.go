package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/auth/authorize", app.authorize)
	r.HandleFunc("/api/auth/token", app.token)
	return r
}
