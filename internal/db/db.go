package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ErrMissingEnv string

func (e ErrMissingEnv) Error() string {
	return "missing required env var: " + string(e)
}

func Connect() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if user == "" || password == "" {
		return nil, ErrMissingEnv("DB_USER, DB_PASSWORD")
	}

	var dsn string

	// Cloud Run → Unix socket
	if os.Getenv("K_SERVICE") != "" {
		instance := os.Getenv("CLOUD_SQL_INSTANCE")
		if instance == "" {
			return nil, ErrMissingEnv("CLOUD_SQL_INSTANCE")
		}

		dsn = fmt.Sprintf(
			"postgres://%s:%s@/%s?host=/cloudsql/%s&sslmode=disable",
			url.QueryEscape(user),
			url.QueryEscape(password),
			user,
			instance,
		)
	} else {
		// Local dev → TCP (Cloud SQL Auth Proxy)
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			url.QueryEscape(user),
			url.QueryEscape(password),
			"127.0.0.1",
			"5432",
			user,
		)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
