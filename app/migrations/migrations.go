package migrations

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // Driver for SQLite3
	_ "github.com/golang-migrate/migrate/v4/source/file"    // Source for file system migrations
)

// MigrateUp performs database migrations to the latest version.
// dsn: The database connection string (e.g., "file:test.db?cache=shared").
// migrationsDir: The directory containing migration files (e.g., "./dbscripts").
func MigrateUp(dsn string, migrationsDir string) error {
	// Check if the migrations directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Printf("Migration directory does not exist: %s", migrationsDir)
		return err
	}

	// Construct the migration source URL
	// Use absolute path for robustness
	absMigrationsDir, err := filepath.Abs(migrationsDir)
	if err != nil {
		return err
	}
	migrateSource := "file://" + absMigrationsDir

	// Construct the database URL
	dbURL := "sqlite3://" + dsn

	// Create a new migrate instance
	m, err := migrate.New(migrateSource, dbURL)
	if err != nil {
		return err
	}

	// Perform the migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		// Handle dirty database state
		if err.Error() == "Dirty database version 1. Fix and force version." {
			log.Printf("Attempting to force version 0 for dirty database.")
			// Force version 0 to clean up dirty state, then retry Up
			if forceErr := m.Force(0); forceErr != nil {
				log.Printf("Failed to force version 0: %v", forceErr)
				return forceErr
			}
			log.Printf("Retrying migration after force version 0.")
			err = m.Up()
			if err != nil && err != migrate.ErrNoChange {
				return err
			}
		}
		return err
	}

	if err == migrate.ErrNoChange {
		log.Println("Database migration: no change")
	} else if err == nil {
		version, dirty, _ := m.Version()
		log.Printf("Database migration successful, current version: %d (dirty: %t)", version, dirty)
	}

	return nil
}
