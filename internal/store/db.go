package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"mbplayer/config"
)

// DBStore is a database store.
type DBStore struct {
	DB *sql.DB
}

// InitDB initializes the database connection.
func InitDB(cfg *config.Config) (*DBStore, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Set the schema search path
	if _, err := db.Exec("SET search_path TO mpdb, public"); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DBStore{DB: db}, nil
}
