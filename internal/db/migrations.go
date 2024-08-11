package db

import (
	"database/sql"
	"log"
	"os"
)

func RunMigrations(db *sql.DB) error {
	migrationScript, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read migration script: %v", err)
		return err
	}

	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		return err
	}
	return nil
}
