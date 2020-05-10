package entity

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserId  uint
	Image   string
	Comment string
}
