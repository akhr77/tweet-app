package entity

import "time"

type Post struct {
	id        uint
	UserId    uint
	Image     string
	Comment   string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
