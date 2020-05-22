package service

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/akhr77/favpic/app/db"
	"github.com/akhr77/favpic/app/entity"
	"github.com/akhr77/favpic/app/utills"
)

type userPost struct {
	ID          int    `db:"id"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	Avater      string `db:"avater"`
	UserProfile string `db:"user_profile"`
	Image       string `db:"image"`
}

type userPosts []userPost

// PostList action: GET /post
// 写真全件を取得
func PostList(w http.ResponseWriter, r *http.Request) {
	var userPosts userPosts
	db := db.GetDB()
	db.Table("users").Select("users.id, users.username, users.email, users.avater, posts.image").Joins("left join posts on posts.user_id = users.id").Scan(&userPosts)

	json.NewEncoder(w).Encode(userPosts)
}

// Upload action: Post /image
// 写真を投稿
func Upload(w http.ResponseWriter, r *http.Request) {
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

	db := db.GetDB()
	var p entity.Post
	p.UserId = 1
	p.Image = "images/develop/" + fileName
	p.Comment = "GORMでのアップロード"
	db.Create(&p)
}

type DownloadImage struct {
	Image string `json:"image"`
}

// Download action: GET /image
// 写真を取得
func Download(w http.ResponseWriter, r *http.Request) {
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
}

// UserbyId action: GET /user
// Userを取得
func UserbyID(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	var u userPost
	id := r.URL.Query().Get("id")
	db.Where("id = ?", id)
	db.Table("users").Select("users.id, users.username, users.email, users.avater, users.user_profile, posts.image").Joins("left join posts on posts.user_id = users.id").Scan(&u)
	imagePath := r.URL.Query().Get("avater")
	// S3への接続
	var awsS3 *utills.AwsS3
	awsS3 = utills.NewAwsS3()
	base64Image, err := awsS3.DownloadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image = "data:image/jpeg;base64," + base64Image
	u.Image = base64Image

	json.NewEncoder(w).Encode(&u)
}

// Update action: PUT /user
// user情報を更新
func Update(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	r.ParseForm()
	log.Print(r)
	id := r.Form.Get("id")
	fileName := r.Form.Get("fileName")
	fileType := r.Form.Get("fileType")
	username := r.Form.Get("username")
	avater := r.Form.Get("avater")

	log.Print(avater)
	if avater != "" {
		// S3への接続
		var (
			err   error
			awsS3 *utills.AwsS3
			// url   string
		)
		awsS3 = utills.NewAwsS3()
		_, err = awsS3.UploadImage(avater, fileName, fileType)
		if err != nil {
			log.Fatal(err)
		}
	}

	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// now := time.Now().In(jst)
	var u entity.User
	db.Where("id = ?", id)
	u.Username = username
	db.Model(&u).Updates(map[string]interface{}{"username": username})
}
