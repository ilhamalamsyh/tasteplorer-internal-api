package database

import (
	"context" // Import context package
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

// Initialize initializes the database connection pool
func Initialize() {
	// Get the connection string from the environment variables
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println("dsn: ", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Parse the connection string
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// Set connection pool configuration (optional)
	config.MaxConns = 10                      // Maximum number of connections in the pool
	config.MaxConnLifetime = 30 * time.Minute // Maximum lifetime of a connection

	// Create the connection pool with a valid context (context.Background())
	dbPool, err := pgxpool.ConnectConfig(context.Background(), config) // Pass context.Background()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Assign the pool to the global DB variable
	DB = dbPool

	// Ping the database to verify the connection
	if err := DB.Ping(context.Background()); err != nil { // Pass context.Background()
		log.Fatalf("Unable to ping the database: %v", err)
	}

	// Log successful database connection
	log.Println("Successfully connected to the database")
}

// Close shuts down the database connection pool
func Close() {
	if DB != nil {
		DB.Close()
	}
}
