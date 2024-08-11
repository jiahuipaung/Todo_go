package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) error {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 验证连接
	if err := DB.Ping(); err != nil {
		return err
	}

	return nil
}

// GetDB 返回数据库连接
func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database not initialized")
	}
	return DB
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
