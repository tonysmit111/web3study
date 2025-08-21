package model

import (
	"github.com/blog/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func init() {
	db := config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) Regist() (err error) {
	db := config.GetDB()
	return db.Create(&u).Error
}

func (u *User) Verify(db *gorm.DB) (err error) {
	return nil
}

func GetUserByName(name string) (User, error) {
	db := config.GetDB()
	u := User{}
	err := db.Where("user_name=?", name).First(&u).Error
	return u, err
}
