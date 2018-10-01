package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "caninodev"
	dbpass = "QpPkW4jkgLc1"
	dbname = "lecheleccion"
)

type Article struct {
	ID                  int
	URL                 string
	UrlToImage          string
	Source              string
	PublicationDate     time.Time
	Title               string
	Body                string
	ExternalReferenceID string
	createdAt           time.Time
	UpdatedAt           time.Time
}

type Collection struct {
	Index []Article
}

func main() {
	initDb()
	http.HandleFunc("/api/articles", articlesHandler)
	//http.HandleFunc("/api/users", usersHandler)
	defer db.Close()

	log.Fatal(http.ListenAndServe("localhost:3001", nil))
}

func initDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbhost, dbport,
		dbuser, dbpass, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	//var (
	//	ID, external_reference_id int
	//	URL, UrlToImage, Source, Body, Title, created_at, updated_ad string
	//
	//
	collection := Collection{}

	err := queryArticles(&collection)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(collection)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func queryArticles(collection *Collection) error {
	rows, err := db.Query(`SELECT t.* FROM collections.articles t LIMIT 50`)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		article := Article{}
		err = rows.Scan(
			&article.ID,
			&article.URL,
			&article.UrlToImage,
			&article.Source,
			&article.PublicationDate,
			&article.Title,
			&article.Body,
			&article.ExternalReferenceID,
			&article.createdAt,
			&article.UpdatedAt)
		if err != nil {
			return err
		}

		collection.Index = append(collection.Index, article)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}


