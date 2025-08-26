package model

import (
	"log"

	"github.com/blog/config"
	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title   string `gorm:"not null"`
    Content string `gorm:"not null"`
    UserID  uint
	User    User `json:"-"`
}

var db *gorm.DB = config.GetDB()

func init() {
	db.AutoMigrate(&Post{})
}

func (p *Post)CreatePost() error{
	log.Println(p)
	return db.Create(p).Error
}

func (p *Post)UpdatePost() error{
	return db.Updates(p).Error
}

func (p Post)DeletePost() error{
	return db.Delete(&p).Error
}

func GetPost(id int) (*Post, error) {
	p:=Post{}
	err := db.Preload("User").First(&p, id).Error
	return &p,err
}

func SelectPosts(pp *Post) (pl []Post, err error) {
	tx := db
	if pp.Title != "" {
		tx = db.Where("title=?", pp.Title)
	}
	if pp.Content != "" {
		tx = tx.Where("content=?", pp.Content)
	}
	pl = []Post{}
	err = tx.Find(&pl).Error
	return
}


