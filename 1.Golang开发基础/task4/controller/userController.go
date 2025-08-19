package controller

import (
	"time"

	"github.com/blog/model"
	"github.com/gin-gonic/gin"
)

func Regist() func(c *gin.Context) {
	return func(c *gin.Context) {
		u := model.User{
			UserName:  "tom",
			Password:  "123456",
			Email:     "123@abc.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		u.Regist()
		c.String(200, "user regist success")
	}
}
