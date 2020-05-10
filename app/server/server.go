package server

import (
	"net/http"

	"github.com/akhr77/favpic/app/service"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/post", service.PostList).Methods(http.MethodGet)
	r.HandleFunc("/image", service.Upload).Methods(http.MethodPost)
	r.HandleFunc("/image", service.Download).Methods(http.MethodGet)
	r.HandleFunc("/user", service.UserbyID).Methods(http.MethodGet)
	r.HandleFunc("/user", service.Update).Methods(http.MethodPost)
	http.ListenAndServe(":8082", r)
}