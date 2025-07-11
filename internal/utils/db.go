package utils

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	adminUser := os.Getenv("ADMIN_USER")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUser == "" || adminPassword == "" {
		log.Fatal("variables ADMIN_USER and/or ADMIN_PASSWORD not set.")
	}

	var err error
	DB, err = sql.Open("sqlite", "./inventory_pot.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	_, err = DB.Exec("INSERT OR IGNORE INTO users (username, password) VALUES (?, ?)", adminUser, adminPassword)
	if err != nil {
		log.Fatal("Error creating admin user: ", err)
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
