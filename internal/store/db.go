package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the pq driver anonymously; it initializes itself.
)

// Store holds the database connection pool.
type Store struct {
	db *sql.DB
}

// NewStore creates a new Store instance and opens the database connection.
func NewStore(dsn string) (*Store, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Verify that the database is reachable.
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection with database: %v", err)
	}

	return &Store{db: db}, nil
}

// Close closes the database connection.
func (s *Store) Close() error {
	return s.db.Close()
}
