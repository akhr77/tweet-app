package main

import (
	"net/http"

	"html/template"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	driverName     = "mysql"
	dataSourceName = "tweet-db:tweet-db@tcp(tweet-db:3306)/tweet-db"
)

type Tweet struct {
	ID    int    `db:"id"`
	Tweet string `db:"tweet"`
}

func init() {
	log.SetPrefix("[Info]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)

}

func createPostTweet(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	v := r.FormValue("tweet")
	_, err = db.NamedExec(`INSERT INTO tweets (tweet) VALUES (:tweet)`, map[string]interface{}{
		"id":    1,
		"tweet": v,
	})
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tweet/", createPostTweet)
	http.ListenAndServe(":8082", nil)
}
