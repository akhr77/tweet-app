package server

import (
	"github.com/akhr77/favpic/app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/post", service.List).Methods(http.MethodGet)
	r.HandleFunc("/post", service.Create).Methods(http.MethodPost)
	r.HandleFunc("/image", service.Show).Methods(http.MethodGet)
	// r.HandleFunc("/user", User)
	// r.HandleFunc("/updateUser", UpdateUser)
	http.ListenAndServe(":8082", r)
}
