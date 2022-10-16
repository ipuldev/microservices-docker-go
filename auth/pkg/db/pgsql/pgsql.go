package pgsql

import (
	"context"

	"github.com/briankliwon/microservices-docker-go/auth/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthModel struct {
	C *pgxpool.Pool
}

func Connect() (*pgxpool.Pool, error) {
	uri := "postgres://root:root@pgsql_auth:5432/auth"
	conn, err := pgxpool.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), `set search_path='public'`)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (db *AuthModel) Insert(auth models.Auth) (*models.Auth, error) {
	err := db.C.QueryRow(context.Background(), "INSERT INTO users (username,password,email) VALUES($1,$2,$3) returning id::text", auth.Username, auth.Password, auth.Email).Scan(&auth.ID)
	if err != nil {
		return nil, err
	}
	return &auth, err
}

func (db *AuthModel) Select(auth models.Auth) (*models.Auth, error) {
	var userData models.Auth
	err := db.C.QueryRow(context.Background(), "SELECT id::text,username,email,password FROM users WHERE username = $1 limit 1", auth.Username).Scan(&userData.ID, &userData.Username, &userData.Email, &userData.Password)
	if err != nil {
		return nil, err
	}
	return &userData, err
}
