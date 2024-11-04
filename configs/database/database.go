package database

import (
	"auth_authorization/configs"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

var (
	once sync.Once
	db   *sqlx.DB
)

// initializeDB initializes the database connection and sets up the connection pool
func initializeDB() {
	var errDB error

	dsn := fmt.Sprintf(
		`user=%s password=%s dbname=%s sslmode=%s`,
		configs.PostgresUsername, "123456", configs.PostgresDBName, configs.PostgresSSLMode,
	)
	db, errDB = sqlx.Connect("postgres", dsn)
	if errDB != nil {
		log.Fatalln("Error connecting to the database:", errDB)
	}

	// Verify the connection to the database is still alive
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxIdleConns(10)   // Set maximum idle connections
	db.SetMaxOpenConns(10)   // Set maximum open connections (including idle)
	db.SetConnMaxLifetime(0) // Unlimited connection lifetime (0 means no limit)

	// Perform auto-migration for (create tables)
	if err := runMigrations(db); err != nil {
		log.Fatalf("Error auto migrating tables: %v", err)
	}
}

// GetDB returns the database connection instance, initializing it if necessary
func GetDB() *sqlx.DB {
	once.Do(initializeDB)
	return db
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{
		MigrationsTable: "migrations", // Custom table for migrations
	})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Path to migration files
		"postgres",          // Database driver
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	err = m.Up() // Apply all pending migrations
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations ran successfully")
	return nil
}
