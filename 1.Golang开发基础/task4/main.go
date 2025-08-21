package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/blog/controller"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	uri := c.Request.RequestURI
	fmt.Println("接收到请求：", uri)
	if strings.HasPrefix(uri, "/user") {
		c.Next()
		return
	}
	tokenString := c.GetHeader("Authorization")
	fmt.Println("token:", tokenString)
	arr := strings.Split(tokenString, " ")
	_token, err := jwt.Parse(arr[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return "abcxyz", nil
	})
	if err != nil || !_token.Valid {
		fmt.Println("token无效", err)
		panic("token 无效，请重新登录")
	}

	if claims, ok := _token.Claims.(jwt.MapClaims); ok {
		c.Set("userId", claims["id"])
		c.Set("userName", claims["userName"])
	}
	c.Next()
	fmt.Println("处理完成")
}

func ErrorGlobalProcess(c *gin.Context) {
	defer func() {
		if r:= recover();r!=nil {
			fmt.Println("recover from panic:", r)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": r,
			})
		}
	}()
	c.Next()
}

func main() {
	r := gin.Default()
	// r.Use(func(c *gin.Context) {
	// 	url := c.Request.URL.Path
	// 	if strings.HasPrefix(url, "/user") {
	// 		c.Next()
	// 		return
	// 	}
	// 	Authentication(c)
	// })
	r.Use(ErrorGlobalProcess, Authentication)
	userGroup := r.Group("/user")
	userGroup.POST("/regist", controller.Regist)
	userGroup.POST("/login", controller.Login)
	r.POST("/add", func(c *gin.Context) {
		v,_:=c.Get("userName")
		s,_:=v.(string)
		c.String(http.StatusOK, s)
	})
	r.Run()

}
