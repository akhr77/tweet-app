package main

import (
	"fmt"
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

type Tweetlist []Tweet

func init() {
	log.SetPrefix("[Info]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tweetlist Tweetlist
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	var Tweet Tweet
	rows, err := db.Queryx("SELECT * FROM tweets")
	for rows.Next() {
		err := rows.StructScan(&Tweet)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB=", Tweet.ID)
		fmt.Println("DB=", Tweet.Tweet)
		tweetlist = append(tweetlist, Tweet)
	}
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("HTML返却=", Tweet)
	t.Execute(w, tweetlist)

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
