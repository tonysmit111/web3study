package main

import (
	"github.com/blog/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/regist", controller.Regist())
	r.Run()
}