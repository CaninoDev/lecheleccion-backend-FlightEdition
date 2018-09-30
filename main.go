package main

import (
	"database/sql"
	"log"
	"net/http"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = 5432
	dbuser = "caninodev"
	dbpass = "80LgjYjvq6zcztiz"
	dbname = "lecheleccion-api_development"
)

func main() {
	initDb()
	defer db.Close()
	http.HandleFunc("/api/articles", articlesHandler)
	http.HandleFunc("/api/users", usersHandler)
	log.Fatal(http.ListenAndServe("localhost:3001", nil))
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func initDB() {
	config := dbConfig()
	var err error
	psqlInfov := fmt.Sprint("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
	dbhost, dbport,
	dbuser, dbpass, dbname)

	db, err = sql.Open("postgres", psqlInfov)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}


