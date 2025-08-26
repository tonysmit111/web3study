package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `json:"-"`
	PostID  uint
	Post    Post `json:"-"`
}

func init() {
	db.AutoMigrate(&Comment{})
}

func (c *Comment) AddCommment() error {
	return db.Create(c).Error
}

func SelectComments(pid uint) ([]Comment, error) {
	cs := []Comment{}
	err := db.Model(&Comment{}).
		Where("post_id=?", pid).
		Find(&cs).Error
	return cs, err
}



