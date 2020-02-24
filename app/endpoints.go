package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"log"

	"github.com/akhr77/favpic/app/infrastructure"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type api struct {
	db *sqlx.DB
}

type userPost struct {
	ID       int    `db:"id"`
	UserName string `db:"username"`
	Email    string `db:"email"`
	Avater   string `db:"avater"`
	Image    string `db:"image"`
}

type userPosts []userPost

func (a *api) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var userPosts userPosts
	var userPost userPost

	rows, err := a.db.Queryx("SELECT u.id,u.username,u.email,u.avater,p.image FROM users u INNER JOIN posts p ON u.id = p.user_id ")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&userPost)
		if err != nil {
			log.Fatal(err)
		}
		userPosts = append(userPosts, userPost)
	}
	json.NewEncoder(w).Encode(userPosts)
}

func (a *api) UploadS3(w http.ResponseWriter, r *http.Request) {
	// formデータの解析
	r.ParseForm()
	file := r.Form.Get("image")
	// data, _ := base64.StdEncoding.DecodeString(image)
	// wb := new(bytes.Buffer)
	// wb.Write(data)

	// S3への接続
	var (
		err   error
		awsS3 *infrastructure.AwsS3
		url   string
	)
	awsS3 = infrastructure.NewAwsS3()
	url, err = awsS3.UploadImage(file, "test", "jpeg")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	log.Println(url)

	http.Redirect(w, r, "/", http.StatusFound)
}

func CreatePostTweet(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	v := r.FormValue("tweet")
	_, err = db.NamedExec(`INSERT INTO tweets (tweet) VALUES (:tweet)`, map[string]interface{}{
		"tweet": v,
	})
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	v := r.FormValue("tweet_delete")
	_, err = db.NamedExec(`DELETE FROM tweets WHERE id = :id`, map[string]interface{}{
		"id": v,
	})
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
