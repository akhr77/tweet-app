package main

import (
	"encoding/json"
	"net/http"

	"app/utills"
	"log"
	"time"

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
	fileName := r.Form.Get("fileName")
	fileType := r.Form.Get("fileType")
	file := r.Form.Get("image")

	// S3への接続
	var (
		err   error
		awsS3 *utills.AwsS3
		// url   string
	)
	awsS3 = utills.NewAwsS3()
	_, err = awsS3.UploadImage(file, fileName, fileType)
	if err != nil {
		log.Fatal(err)
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	_, err = a.db.NamedExec(`INSERT INTO posts (user_id,image,comment,created_at,updated_at) VALUES (:userId,:image,:comment,:created,:updated)`, map[string]interface{}{
		"userId":  1,
		"image":   "images/develop/" + fileName,
		"comment": "アップロードできたー",
		"created": now,
		"updated": now,
	})
	if err != nil {
		log.Fatal(err)
	}

	// http.Redirect(w, r, "/", http.StatusFound)
}

type DownloadImage struct {
	Image string `json:"image"`
}

func (a *api) DownloadS3(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータのパース
	imagePath := r.URL.Query().Get("image")
	// S3への接続
	var awsS3 *utills.AwsS3
	awsS3 = utills.NewAwsS3()
	base64Image, err := awsS3.DownloadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image = "data:image/jpeg;base64," + base64Image
	downloadImage := DownloadImage{Image: base64Image}

	json.NewEncoder(w).Encode(downloadImage)
	// fmt.Fprintln(w, base64Image)
}

// func DeleteTweet(w http.ResponseWriter, r *http.Request) {
// 	db, err := sqlx.Open(driverName, dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	v := r.FormValue("tweet_delete")
// 	_, err = db.NamedExec(`DELETE FROM tweets WHERE id = :id`, map[string]interface{}{
// 		"id": v,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	http.Redirect(w, r, "/", http.StatusFound)
// }
