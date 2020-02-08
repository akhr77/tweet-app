package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	dsn        = "favpic:favpic@tcp(favpicDB:3306)/favpic?parseTime=true"
)

func init() {
	log.SetPrefix("[Log]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := &api{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/", app.IndexHandler)
	r.HandleFunc("/tweet/", CreatePostTweet)
	r.HandleFunc("/tweet_delete", DeleteTweet)
	http.ListenAndServe(":8082", r)
}
