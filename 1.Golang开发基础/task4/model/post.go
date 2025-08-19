package model

import (
	"github.com/blog/config"
	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title   string `gorm:"not null"`
    Content string `gorm:"not null"`
    UserID  uint
    User    User
}

func init() {
	db:=config.GetDB()
	db.AutoMigrate(&Post{})
}