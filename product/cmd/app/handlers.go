package main

import (
	"encoding/json"
	"net/http"

	"github.com/briankliwon/microservices-docker-go/product/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	product, err := app.product.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(product)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Product have been listed")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.Product
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	insertResult, err := app.product.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}
	response := &models.InsertResponse{
		Message:   "Insert Data success",
		Insert_id: insertResult.InsertedID.(primitive.ObjectID).Hex(),
	}

	b, err := json.Marshal(response)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New product have been created, id=%s", insertResult.InsertedID)
	app.infoLog.Println("Product have been listed")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.product.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d product(s)", deleteResult.DeletedCount)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.product.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Movie not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a product")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
