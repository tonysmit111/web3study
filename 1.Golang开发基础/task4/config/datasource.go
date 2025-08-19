package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var db *gorm.DB

func init() {
	_db, err := gorm.Open(sqlite.Open("blog.db"))
	if err != nil {
		panic(err)
	}
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
