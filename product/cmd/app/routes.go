package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/product/", app.all).Methods("GET")
	r.HandleFunc("/api/product/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/api/product/", app.insert).Methods("POST")
	r.HandleFunc("/api/product/{id}", app.delete).Methods("DELETE")

	return r
}
