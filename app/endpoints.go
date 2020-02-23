package main

import (
	"encoding/json"
	"net/http"

	"log"

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
	// sess := session.Must(session.NewSession(&aws.Config{
	// 	Region:      aws.String("ap-northeast-1"),
	// 	Credentials: credentials.NewSharedCredentials("", "defult"),
	// }))

	// uploader := s3manager.NewUploader(sess)

	// result, err := uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String("amplify-favpic-favpic-145648-deployment"),
	// 	Key:    aws.String("favpic"),
	// 	Body:   data,
	// })

	// //結果表示
	// if err == nil {
	// 	fmt.Println(result.Location)
	// } else {
	// 	fmt.Println("error happend!!!")
	// }
	r.ParseForm()
	message := r.Form.Get("message")
	image := r.Form.Get("image")
	log.Println(message)
	log.Println(image)

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
