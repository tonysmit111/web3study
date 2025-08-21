package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blog/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Regist(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = string(hashedPassword)

	if err := user.Regist(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfuly",
	})

}

func Login(c *gin.Context) {
	var user model.User
	if err:= c.ShouldBindJSON(&user);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(&user)

	var storedUser model.User
	var err error
	storedUser, err = model.GetUserByName(user.UserName)
	if err!=nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Invalid username or password",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": storedUser.ID,
		"userName" : storedUser.UserName,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(GetSecrketKey()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to generate token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token":tokenString,
	})
}

func GetSecrketKey() string {
	return "abcxyz"
}
