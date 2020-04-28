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
	r.HandleFunc("/upload/s3", app.UploadS3)
	r.HandleFunc("/download/s3", app.DownloadS3)
	r.HandleFunc("/user", app.user)
	r.HandleFunc("/updateUser", app.UpdateUser)
	http.ListenAndServe(":8082", r)
}
