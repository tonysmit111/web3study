package model

import (
	"github.com/blog/config"
	"gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content string `gorm:"not null"`
    UserID  uint
    User    User
    PostID  uint
    Post    Post
}

func init() {
	db:=config.GetDB()
	db.AutoMigrate(&Comment{})
}
