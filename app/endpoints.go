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

type user struct {
	ID               int            `db:"id"`
	UserName         string         `db:"username"`
	Email            string         `db:"email"`
	EncrypedPassword sql.NullString `db:"encryped_password"`
	Avater           string         `db:"avater"`
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

type users []user

func (a *api) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var users users
	var user user

	rows, err := a.db.Queryx("SELECT id,username,email,encryped_password,avater FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
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
