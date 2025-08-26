package controller

import (
	"net/http"
	"strconv"

	"github.com/blog/model"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// p := model.Post{}
	var p model.Post
	c.ShouldBindJSON(&p)
	uid, exi := c.Get("userId")
	if exi {
		p.UserID = uint(uid.(float64))
	}

	if err := p.CreatePost(); err != nil {
		panic(err)
	}
	// if  err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{
	// 	// 	"error": err,
	// 	// })
	// 	// return
	// 	panic(err)
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "文章创建成功",
	})
}

func UpdatePost(c *gin.Context) {
	p := model.Post{}
	c.ShouldBindJSON(&p)
	if p.ID <= 0 {
		panic("要更新的文章id不能为空")
	}
	uid, exi := c.Get("userId")
	if exi {
		p.UserID = uint(uid.(float64))
	}
	ownp, err := model.GetPost(int(p.ID))
	if err != nil {
		panic(err.Error())
	}
	if ownp.User.ID != uid {
		panic("不是当前文章的作者，不能更新这篇文章")
	}

	if err := p.UpdatePost(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
	})
}

func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		panic(err)
	}
	p, err := model.GetPost(id)
	if err != nil {
		panic(err)
	}
	uid, exi := c.Get("userId")
	if !exi || p.UserID != uint(uid.(float64)) {
		panic("不是当前文章的作者，不能删除这篇文章")
	}
	if err := p.DeletePost(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "删除成功",
	})

}

func SelectPosts(c *gin.Context) {
	p := model.Post{}
	c.ShouldBindJSON(&p)
	ps, err := model.SelectPosts(&p)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"list": ps,
	})
}

func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		panic(err)
	}
	p, err := model.GetPost(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"obj": p,
	})
}
