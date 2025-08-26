package controller

import (
	"net/http"
	"strconv"

	"github.com/blog/model"
	"github.com/gin-gonic/gin"
)

func AddCommment(c *gin.Context) {
	cmt := model.Comment{}
	c.ShouldBindBodyWithJSON(&cmt)
	uid, exi := c.Get("userId")
	if exi {
		cmt.UserID = uint(uid.(float64))
	}
	_, err := model.GetPost(int(cmt.PostID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := cmt.AddCommment(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加评论成功",
	})
}

func SelectComments(c *gin.Context) {
	postId, err := strconv.Atoi(c.Query("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	cs, err := model.SelectComments(uint(postId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H {
		"list":cs,
	})

}
