package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL, -- Storing as BLOB for byte slice
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

func main(){
	dbName := "data.db"

	_ = os.Remove(dbName)

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	defer func(){
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
	_, err = db.Exec(schema)
	if err != nil{
		log.Fatal("failed to create schema: ", err)
	}

	fmt.Println("Database schema created successfully")
}
