package main

import (
	"database/sql"
	"freecodecamp/database/repository_stuff/repository"

	//"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func main() {

	dbName := "users.db"

	db, err := connectToDatabase(dbName)
	checkErr(err)

	defer db.Close()

	fmt.Println("database connection established")

	repo := repository.NewSQLUserRepository(db)

	printUsers(repo)

}

func printUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()
	checkErr(err)
	for _, user := range users {
		fmt.Println("id: ", user.ID, "name: ", user.Name, "email: ", user.Email, "created_at: ", user.CreatedAt)
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

func checkErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}





