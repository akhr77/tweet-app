package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avater      string `json:"avater"`
	UserProfile string `json:"userProfile"`
	GuestFlag   string `json:"guestFlag"`
	AdminFlag   string `json:"adminFlag"`
}
