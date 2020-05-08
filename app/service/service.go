package service

mport (
	"encoding/json"
	"net/http"

	"github.com/akhr77/favpic/app/utills"
	"log"
	"time"

)

type userPost struct {
	ID       int    `db:"id"`
	UserName string `db:"username"`
	Email    string `db:"email"`
	Avater   string `db:"avater"`
	Image    string `db:"image"`
}

type userPosts []userPost

func List(w http.ResponseWriter, r *http.Request) {
	log.Print("Listにはきたよ")
	var userPosts userPosts
	db := db.GetDB()
	db.Table("users").Select("users.id, users.username, users.email, users.avater").Joins("left join posts on users.id = posts.user_id").Scan(&userPosts)

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
}

type DownloadImage struct {
	Image string `json:"image"`
}

func Show(w http.ResponseWriter, r *http.Request) {
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

func (a *api) User(w http.ResponseWriter, r *http.Request) {
	var user user
	id := r.URL.Query().Get("id")
	imagePath := r.URL.Query().Get("avater")
	// S3への接続
	var awsS3 *utills.AwsS3
	awsS3 = utills.NewAwsS3()
	base64Image, err := awsS3.DownloadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image = "data:image/jpeg;base64," + base64Image
	user.Image = base64Image

	rows, err := a.db.NamedQuery(`SELECT u.id,u.username,u.avater,u.user_profile FROM users u WHERE u.id = :userId`, map[string]interface{}{"userId": id})
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode(user)
}

func (a *api) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// formデータの解析
	r.ParseForm()
	log.Print(r)
	userId := r.Form.Get("userId")
	fileName := r.Form.Get("fileName")
	fileType := r.Form.Get("fileType")
	userName := r.Form.Get("userName")
	avater := r.Form.Get("avater")
	userProfile := r.Form.Get("userProfile")

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

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	rows, err := a.db.NamedExec(`UPDATE users SET username = :username WHERE id = :userid`, map[string]interface{}{
		"userid":   userId,
		"username": userName,
		// "avater":       "images/develop/" + fileName,
		"user_profile": userProfile,
		"updated":      now,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(rows)
}
