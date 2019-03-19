package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/api/articles", GetArticles).Methods("GET")
	router.HandleFunc("/api/article/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/api/bias/{id}", GetBias).Methods("GET")
	router.HandleFunc("/api/user/{id}", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))
}

func GetArticles(w http.ResponseWriter, r *http.Request) {

	// ...
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	// ...
}

func GetBias(w http.ResponseWriter, r *http.Request) {
	// ...
}

func GetUser(w http.ResponseWriter, r *http.Request) {
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


