package pgsql

import (
	"context"

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
