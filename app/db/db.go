package db

import (
	"log"

	"github.com/akhr77/favpic/app/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init DB初期設定
func Init() {
	DBMS := "mysql"
	CONNECT := "favpic:favpic@tcp(favpicDB:3306)/favpic?parseTime=true"
	_db, err := gorm.Open(DBMS, CONNECT)
	db = _db
	if err != nil {
		log.Fatal(err)
	}

	// autoMigration()
}

// GetDB DBアクセサ取得
func GetDB() *gorm.DB {
	return db
}

// Close DB接続終了
func Close() {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User)
	db.AutoMigrate(&entity.Post)
}
