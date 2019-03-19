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
	dbpass = "testing"
	dbname = "lecheleccion"
)

func main() {
	initConnDB()
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

func initConnDB() {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}


