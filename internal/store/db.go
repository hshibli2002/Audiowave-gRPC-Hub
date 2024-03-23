package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
	"mbplayer/config"     // Adjust the import path according to your project structure
)

// DBStore is a database store that holds a SQL DB connection.
type DBStore struct {
	DB *sql.DB
}

// InitDB initializes the database connection using configuration from the cfg parameter.
func InitDB(cfg *config.Config) (*DBStore, error) {
	// Ensure that sslmode is set to 'require' as per your Neon connection requirements
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Optionally, set the schema search path to mpdb, public if necessary for your application
	if _, err := db.Exec("SET search_path TO mpdb, public"); err != nil {
		_ = db.Close() // Attempt to close the database connection on error
		return nil, fmt.Errorf("failed to set search path: %w", err)
	}

	// Verify the database connection
	if err := db.Ping(); err != nil {
		_ = db.Close() // Attempt to close the database connection on error
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Return the DBStore containing the established database connection
	return &DBStore{DB: db}, nil
}
