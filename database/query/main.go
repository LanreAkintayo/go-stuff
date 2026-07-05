package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

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


	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal into json

	bytes, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal("failed to marshal users: ", err)
	}

	
	fmt.Println(string(bytes))

	// fmt.Println("Creating database schema...")
	// createTable(db)
	// fmt.Println("Table has been created")

	// lastId, err := createUser(db, "a", "a@gmail.com", "password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Last user id is, ", lastId)
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

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
