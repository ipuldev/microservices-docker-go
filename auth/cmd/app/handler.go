package main

import (
	"encoding/json"
	"net/http"

	"github.com/briankliwon/microservices-product-catalog/auth/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) authorize(w http.ResponseWriter, r *http.Request) {
	_, err := app.srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpResponse := &models.HttpResponse{
		Message: "Authentiicated",
	}
	b, err := json.Marshal(httpResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) token(w http.ResponseWriter, r *http.Request) {
	app.srv.HandleTokenRequest(w, r)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	var m models.Auth
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	if err != nil {
		app.serverError(w, err)
	}
	m.Password = string(passwordEncrypted)
	insertResult, err := app.auth.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.ID)

	httpResponse := &models.HttpResponse{
		Message: "New user have been created",
		OauthData: models.Oauth2Key{
			ClientID:     app.oauth2.ClientID,
			ClientSecret: app.oauth2.ClientSecret,
		},
	}
	b, err := json.Marshal(httpResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var m models.Auth
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userData, err := app.auth.Select(m)
	if err != nil {
		app.serverError(w, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(m.Password))
	if err != nil {
		app.serverError(w, err)
	}

	httpResponse := &models.HttpResponse{
		Message: "Login success",
		OauthData: models.Oauth2Key{
			ClientID:     app.oauth2.ClientID,
			ClientSecret: app.oauth2.ClientSecret,
		},
	}
	app.srv.HandleTokenRequest(w, r)
	b, err := json.Marshal(httpResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
