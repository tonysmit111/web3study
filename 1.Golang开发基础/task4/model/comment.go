package model

import (
	"github.com/blog/config"
	"gorm.io/gorm"
)

type Comment struct {
    gorm.Model
    Content string `gorm:"not null"`
    UserID  uint
    PostID  uint
}

func init() {
	db:=config.GetDB()
	db.AutoMigrate(&Comment{})
}
