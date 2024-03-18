package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"mbplayer/config"
	"mbplayer/internal/Handlers"
	"mbplayer/internal/Services"
	"mbplayer/internal/store"
	queries "mbplayer/internal/store/Queries"
	pb "mbplayer/pkg/grpcapi"
	"net"
	"os"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbStore, err := store.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}(dbStore.DB)

	// Insert mock data
	insertMockArtists(dbStore.DB, 20)
	insertMockSongs(dbStore.DB, 200)
	insertMockUsers(dbStore.DB, 100)
	insertMockPlaylists(dbStore.DB, 50)
	insertMockPlaylistSongs(dbStore.DB, 20)
	insertMockUserFollowsArtists(dbStore.DB, 20)
	insertMockUserLikesSongs(dbStore.DB, 20)

	// Initialize services
	songQueries := queries.NewSongQueries(dbStore.DB)
	songService := Services.NewSongService(songQueries)

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register your services with the gRPC server
	pb.RegisterSongsServiceServer(grpcServer, Handlers.NewSongServer(songService))

	// Determine the port to listen on
	port := "50051"
	if envPort := os.Getenv("GRPC_PORT"); envPort != "" {
		port = envPort
	}

	// Start listening on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	fmt.Printf("gRPC server listening on port %s\n", port)

	// Start serving gRPC requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}

func insertMockArtists(db *sql.DB, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO artists (name, bio, created_at, followers_count) VALUES ($1, $2, $3, $4)",
			randomString(10), // Generate a random string for name
			randomString(50), // Generate a random string for bio
			time.Now(),       // Use the current time
			rand.Intn(1000),  // Generate a random number for followers_count
		)
		if err != nil {
			fmt.Println("Error inserting mock artist:", err)
		}
	}
}

func insertMockSongs(db *sql.DB, count int) {
	// Get the number of artists to ensure we have valid foreign key references for artist_id
	artistCount := getCountFromTable(db, "mpdb.artists")

	for i := 0; i < count; i++ {
		// Assuming that not all songs have an artist, some can be null
		var artistId sql.NullInt64
		if rand.Intn(10) < 8 { // 80% chance the song will have an artist
			artistId = sql.NullInt64{Int64: int64(rand.Intn(artistCount) + 1), Valid: true}
		} else {
			artistId = sql.NullInt64{Valid: false}
		}

		duration := rand.Intn(300) // Duration between 0 and 299 seconds
		_, err := db.Exec(
			"INSERT INTO mpdb.songs (title, artist_id, duration, created_at) VALUES ($1, $2, $3, $4)",
			randomString(10), // Generate a random string for the title
			artistId,         // Random or no artist
			duration,         // Random duration
			time.Now(),       // Use the current time for created_at
		)
		if err != nil {
			fmt.Println("Error inserting mock song:", err)
		}
	}
}

func insertMockUsers(db *sql.DB, count int) {
	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO mpdb.users (username, email, password, created_at, followings_count) VALUES ($1, $2, $3, $4, $5)",
			randomString(8),                      // Generate a random string for the username
			fmt.Sprintf("user%d@example.com", i), // Create a mock email
			randomString(12),                     // Generate a random string for the password
			time.Now(),                           // Use the current time
			rand.Intn(100),                       // Generate a random number for followings_count
		)
		if err != nil {
			fmt.Println("Error inserting mock user:", err)
		}
	}
}

func insertMockPlaylists(db *sql.DB, count int) {
	// Get the number of users to ensure we have valid foreign key references
	userCount := getCountFromTable(db, "mpdb.users")
	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO mpdb.playlists (title, creator_id, created_at) VALUES ($1, $2, $3)",
			randomString(15),       // Generate a random string for the title
			rand.Intn(userCount)+1, // Generate a random user_id, ensuring it exists
			time.Now(),             // Use the current time
		)
		if err != nil {
			fmt.Println("Error inserting mock playlist:", err)
		}
	}
}

func insertMockPlaylistSongs(db *sql.DB, count int) {
	// Get the number of playlists and songs to ensure we have valid foreign key references
	playlistCount := getCountFromTable(db, "mpdb.playlists")
	songCount := getCountFromTable(db, "mpdb.songs")

	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO mpdb.playlist_songs (playlist_id, song_id) VALUES ($1, $2)",
			rand.Intn(playlistCount)+1, // Generate a random playlist_id, ensuring it exists
			rand.Intn(songCount)+1,     // Generate a random song_id, ensuring it exists
		)
		if err != nil {
			fmt.Println("Error inserting mock playlist_song:", err)
		}
	}
}

func insertMockUserFollowsArtists(db *sql.DB, count int) {
	// Get the number of users and artists to ensure we have valid foreign key references
	userCount := getCountFromTable(db, "mpdb.users")
	artistCount := getCountFromTable(db, "mpdb.artists")

	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO mpdb.user_follows_artists (user_id, artist_id) VALUES ($1, $2)",
			rand.Intn(userCount)+1,   // Generate a random user_id, ensuring it exists
			rand.Intn(artistCount)+1, // Generate a random artist_id, ensuring it exists
		)
		if err != nil {
			fmt.Println("Error inserting mock user_follows_artist:", err)
		}
	}
}

func insertMockUserLikesSongs(db *sql.DB, count int) {
	// Get the number of users and songs to ensure we have valid foreign key references
	userCount := getCountFromTable(db, "mpdb.users")
	songCount := getCountFromTable(db, "mpdb.songs")

	for i := 0; i < count; i++ {
		_, err := db.Exec(
			"INSERT INTO mpdb.user_likes_songs (user_id, song_id) VALUES ($1, $2)",
			rand.Intn(userCount)+1, // Generate a random user_id, ensuring it exists
			rand.Intn(songCount)+1, // Generate a random song_id, ensuring it exists
		)
		if err != nil {
			fmt.Println("Error inserting mock user_likes_song:", err)
		}
	}
}

// getCountFromTable is a utility function to get the count of records from a table.
func getCountFromTable(db *sql.DB, tableName string) int {
	var count int
	row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))
	if err := row.Scan(&count); err != nil {
		fmt.Printf("Error getting count from table %s: %v\n", tableName, err)
		return 0
	}
	return count
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
