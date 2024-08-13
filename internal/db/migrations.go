package db

import (
	"database/sql"
	"log"
	"os"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	migrationScript, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read migration script: %v", err)
		return err
	}

	// 以分号分割 SQL 脚本
	statements := strings.Split(string(migrationScript), ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if len(stmt) > 0 {
			_, err := db.Exec(stmt)
			if err != nil {
				log.Fatalf("Failed to execute statement: %v", err)
				return err
			}
		}
	}
	return nil
}
