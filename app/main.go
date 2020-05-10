package main

import (
	"log"

	"github.com/akhr77/favpic/app/db"
	"github.com/akhr77/favpic/app/server"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	db.Init()
	server.Init()
	db.Close()
}
