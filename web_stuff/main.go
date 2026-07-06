package main

import (
	"database/sql"
	"fmt"
	"log"
	// "net/http"
	"os"
	_ "github.com/mattn/go-sqlite3"

)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepository
}

func main() {

	db, err := connectToDatabase("users.db")
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo: NewSQLUserRepository(db),
	}

	fmt.Println("Starting server on :8080")
	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
