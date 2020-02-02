package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[Log]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/tweet/", CreatePostTweet)
	r.HandleFunc("/tweet_delete", DeleteTweet)
	http.ListenAndServe(":8082", r)
}
