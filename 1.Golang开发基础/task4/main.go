package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/blog/controller"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func init() {
	f,err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, f))
	defer f.Close()
}

var ignoreAuthUrls []string = []string{"/comment/select","/post/select","/post/get"}

func Authentication(c *gin.Context) {
	uri := c.Request.RequestURI
	if strings.HasPrefix(uri, "/user") {
		c.Next()
		return
	}
	for _,u := range ignoreAuthUrls {
		if u == uri {
			c.Next()
			return
		}
	}
	
	tokenString := c.GetHeader("Authorization")
	log.Println("token:", tokenString)
	arr := strings.Split(tokenString, " ")
	_token, err := jwt.Parse(arr[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("abcxyz"), nil
	})
	if err != nil || !_token.Valid {
		log.Println("token无效", err)
		// panic("token 无效，请重新登录")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error" : "token 无效，请重新登录",
		})
	}

	if claims, ok := _token.Claims.(jwt.MapClaims); ok {
		c.Set("userId", claims["id"])
		c.Set("userName", claims["userName"])
	}
	c.Next()
	
}

func ErrorGlobalProcess(c *gin.Context) {
	uri := c.Request.RequestURI
	log.Println("接收到请求：", uri)
	defer func() {
		if r := recover(); r != nil {
			log.Print("recover from panic:", r)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": r,
			})		
		}
		log.Println("处理完成")
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
// v,_:=c.Get("userName")
// s,_:=v.(string)
	postGroup := r.Group("/post")
	postGroup.POST("/create", controller.CreatePost)
	postGroup.POST("/update", controller.UpdatePost)
	postGroup.POST("/select", controller.SelectPosts)
	postGroup.GET("/get", controller.GetPost)
	postGroup.DELETE("/delete", controller.DeletePost)

	commentGroup := r.Group("comment")
	commentGroup.GET("select", controller.SelectComments)
	commentGroup.POST("add", controller.AddCommment)
	r.Run()

}
