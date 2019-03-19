package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

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

type Article struct {
	ID                  int
	URL                 string
	URLToImage          string
	Source              string
	PublicationDate     time.Time
	Title               string
	Body                string
	ExternalReferenceID int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Bias struct {
	ID           int
	Libertarian  float32
	Green        float32
	Liberal      float32
	Conservative float32
	biasableType string
	biasableID   int
	createdAt    time.Time
	updatedAt    time.Time
}

type ClientMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func main() {
	initConnDB()
	defer db.Close()
	createRouter()
}

func createRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/api/articles", GetArticles).Methods("GET")
	router.HandleFunc("/api/article/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/api/bias/{id}", GetBias).Methods("GET")
	router.HandleFunc("/api/user/{id}", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var msg ClientMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println("Error parsing client request.")
	}
	msg.Type = params["type"]
	msg.Payload = params["payload"]
	if msg.Type != "quantity" {
		log.Println("Malformed request or wrong endpoint")
	}
	queryArticles(w, msg.Payload)

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


