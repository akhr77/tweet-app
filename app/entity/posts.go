package entity

import "time"

type Post struct {
	id        uint
	UserId    uint
	Image     string
	Comment   string
	CreatedAt time.Time `json:"createdAt"`
	updatedAt time.Time `json:"updatedAt"`
}
