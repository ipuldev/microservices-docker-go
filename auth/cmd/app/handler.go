package main

import "net/http"

func (app *application) authorize(w http.ResponseWriter, r *http.Request) {
	err := app.srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (app *application) token(w http.ResponseWriter, r *http.Request) {
	app.srv.HandleTokenRequest(w, r)
}
