package server

import (
	"log"
	"net/http"

	"github.com/akhr77/favpic/app/service"

	"github.com/gorilla/mux"
)

func Init() {
	r := Router()
	http.ListenAndServe(":8082", r)
}

// NewMux create the handler.
func Router() http.Handler {
	r := mux.NewRouter()

	r.Use(requestLogger)

	r.HandleFunc("/post", service.PostList).Methods(http.MethodGet)
	r.HandleFunc("/image", service.Upload).Methods(http.MethodPost)
	r.HandleFunc("/image", service.Download).Methods(http.MethodGet)
	r.HandleFunc("/user", service.UserbyID).Methods(http.MethodGet)
	r.HandleFunc("/user", service.Update).Methods(http.MethodPost)

	return r
}

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", r)
		handler.ServeHTTP(w, r)
	})
}
