package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var db *sql.DB

var addr = flag.String("addr", "localhost:3001", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

}

var conn websocket.Conn

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "caninodev"
	dbpass = "testing"
	dbname = "lecheleccion"
)

// Message corresponds to the client's request for data of any type.
type Message struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// Article is the model Article with all relevant fields
type Article struct {
	ID                  int
	URL                 string
	URLToImage          string
	Source              string
	PublicationDate     time.Time
	Title               string
	Body                string
	ExternalReferenceID string
	createdAt           time.Time
	UpdatedAt           time.Time
}

// Bias corresponds to the aforementioned data in the DB
type Bias struct {
	ID           int
	Libertarian  float32
	Green        float32
	Liberal      float32
	Conservative float32
	biasableType string
	biasableID int
	createdAt time.Time
	updatedAt time.Time
}

// User corresponds to the Model User
type User struct {
	ID   int
	Name string
}

// Group a group of Users
type Group struct {
	Index []User
}

func main() {
	initDb()

	http.HandleFunc("/api/articles", articlesHandler)
	http.HandleFunc("/api/bias", biasHandler)
	http.HandleFunc("/api/users", usersHandler)

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
	//var msgType int
	var msg []byte
	var err error

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer conn.Close()

	for {
		_, msg, err = conn.ReadMessage()
		if err != nil {
			return
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			panic(err)
		}
		switch message.Type {
			 case "articles":
			 queryArticles(*conn, &message.Payload)
			 default:
			 http.Error(w, "Malformed request or wrong endpoint.", 500)
		}
		fmt.Printf("%s sent type: %s payload: %s\n", conn.RemoteAddr(), string(message.Type), string(message.Payload))
	}
}

func biasHandler(w http.ResponseWriter, r *http.Request) {
	//var msgType int
	var msg []byte
	var err error

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer conn.Close()

	for {
		_, msg, err = conn.ReadMessage()
		if err != nil {
			return
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			panic(err)
		}
		switch message.Type {
			 case "bias":
			 queryBias(*conn, &message.Payload)
			 default:
			 http.Error(w, "Malformed request or wrong endpoint.", 500)
		}
		fmt.Printf("%s sent type: %s payload: %s\n", conn.RemoteAddr(), string(message.Type), string(message.Payload))
	}
}

func queryArticles(conn **websocket.Conn, msg *[]byte) error {
	var searchTerms []string
	if msg == nil {
		searchTerms[0] = "t.*"
	} else {
		searchTerms = strings.Split(string(*msg), " ")
	}

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
			&article.URLToImage,
			&article.Source,
			&article.PublicationDate,
			&article.Title,
			&article.Body,
			&article.ExternalReferenceID,
			&article.createdAt,
			&article.UpdatedAt)

		err = (*conn).WriteJSON(&article)

		if err != nil {
			return err
		}
		//fmt.Println(article)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func usersHandler(w http.ResponseWriter, _ *http.Request) {
	group := Group{}

	err := queryUsers(&group)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))

}

func queryUsers(group *Group) error {
	rows, err := db.Query(`SELECT t.* FROM collections.users t`)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Name)
		if err != nil {
			return err
		}

		group.Index = append(group.Index, user)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
