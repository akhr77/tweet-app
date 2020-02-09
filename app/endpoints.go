package main

import (
	"database/sql"
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
	ID               int            `db:"id"`
	UserName         string         `db:"username"`
	Email            string         `db:"email"`
	EncrypedPassword sql.NullString `db:"encryped_password"`
	Avater           string         `db:"avater"`
	Image            string         `db:"image"`
}

// type user struct {
// 	ID                 int            `db:"id"`
// 	UserName           string         `db:"username"`
// 	Email              string         `db:"email"`
// 	EncrypedPassword   sql.NullString `db:"encryped_password"`
// 	Avater             string         `db:"avater"`
// 	ResetPasswordToken string         `db:"reset_password_token"`
// 	ConfirmationToken  string         `db:"confirmation_token"`
// 	GuestFlag          string         `db:"guest_flag"`
// 	AdminFlag          string         `db:"admin_flag"`
// 	CreatedAt          time.Time      `db:"created_at"`
// 	UpdatedAt          time.Time      `db:"updated_at"`
// }

type userPosts []userPost

func (a *api) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var userPosts userPosts
	var userPost userPost

	rows, err := a.db.Queryx("SELECT u.id,u.username,u.email,u.encryped_password,u.avater,p.image FROM users u INNER JOIN posts p ON u.id = p.user_id ")
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
