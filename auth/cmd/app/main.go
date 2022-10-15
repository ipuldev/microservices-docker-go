package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/briankliwon/microservices-product-catalog/auth/pkg/db/pgsql"
	"github.com/briankliwon/microservices-product-catalog/auth/pkg/models"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauth_model "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/google/uuid"
)

func main() {
	//HTTP properties
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")

	// Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	manager := manage.NewDefaultManager()

	//token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	//client memory store
	clientStore := store.NewClientStore()

	clientID := uuid.New().String()[:8]
	clientSecret := uuid.New().String()[:8]
	clientStore.Set(clientID, &oauth_model.Client{
		ID:     clientID,
		Secret: clientSecret,
		Domain: "http://localhost",
	})

	manager.MapClientStorage(clientStore)
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	pgsqlConnection, err := pgsql.Connect()
	if err != nil {
		errLog.Fatal(err)
	}
	infoLog.Printf("Database connection established")
	app := &application{
		infoLog:  infoLog,
		errorLog: errLog,
		srv:      srv,
		auth: &pgsql.AuthModel{
			C: pgsqlConnection,
		},
		oauth2: &models.Oauth2Key{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	http_server := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = http_server.ListenAndServe()
	errLog.Fatal(err)
}
