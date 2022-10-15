package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/briankliwon/microservices-product-catalog/auth/pkg/db/pgsql"
	"github.com/briankliwon/microservices-product-catalog/auth/pkg/models"
	"github.com/go-oauth2/oauth2/v4/server"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	srv      *server.Server
	auth     *pgsql.AuthModel
	oauth2   *models.Oauth2Key
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
