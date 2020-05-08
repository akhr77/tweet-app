package main

import (
	"app/db"
	"app/server"
	"log"
)

// const (
// 	driverName = "mysql"
// 	dsn        = "favpic:favpic@tcp(favpicDB:3306)/favpic?parseTime=true"
// )

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	db.Init()
	server.Init()
	db.Close()
	// db, err := sqlx.Open(driverName, dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// app := &api{db: db}
	// r := mux.NewRouter()
	// r.HandleFunc("/", app.IndexHandler)
	// r.HandleFunc("/upload/s3", app.UploadS3)
	// r.HandleFunc("/download/s3", app.DownloadS3)
	// r.HandleFunc("/user", app.User)
	// r.HandleFunc("/updateUser", app.UpdateUser)
	// http.ListenAndServe(":8082", r)
}
