package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentApi struct{}

var CommentApiApp = new(CommentApi)

func (m *CommentApi) CreateComment(c *gin.Context) {
	comment := &Comment{}
	err := c.ShouldBindJSON(comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}

	comment.UserId = uint(userId.(float64))
	DB.Create(comment)
	c.JSON(http.StatusOK, "评论成功")
}

func (m *CommentApi) CommentList(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Query("postId"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments := []*Comment{}
	DB.Where("post_id = ?", postId).Find(&comments)
	c.JSON(http.StatusOK, comments)
	return
}
