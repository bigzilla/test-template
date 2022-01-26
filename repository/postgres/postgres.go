package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func Open(cfg Config) (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
