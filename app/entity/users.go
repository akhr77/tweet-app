package entity

import "time"

type User struct {
	id          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Avater      string    `json:"avater"`
	UserProfile string    `json:"userProfile"`
	GuestFlag   string    `json:"guestFlag"`
	AdminFlag   string    `json:"adminFlag"`
	CreatedAt   time.Time `json:"createdAt"`
	updatedAt   time.Time `json:"updatedAt"`
}
