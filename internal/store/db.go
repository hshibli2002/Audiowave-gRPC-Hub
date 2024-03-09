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

	//fmt.Println("DSN:", dsn)

	// Open the connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DBStore{DB: db}, nil
}
