package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/briankliwon/microservices-product-catalog/product/pkg/db/mongodb"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	product  *mongodb.ProductModel
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
