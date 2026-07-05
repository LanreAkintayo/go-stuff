package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

func main() {
	dbName := "users.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	defer func() {
		fmt.Println("Closing database connection")
		if err := db.Close(); err != nil {
			log.Fatal("failed to close db: ", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping db: ", err)
	}
	fmt.Println("Database connection established successfully")

	fmt.Println("Creating database schema...")
	createTable(db)
	fmt.Println("Table has been created")

	lastId, err := createUser(db, "Lanre", "lanre@gmail.com", "password")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Last user id is, ", lastId)
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name, email, plainPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hashedPass))
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
