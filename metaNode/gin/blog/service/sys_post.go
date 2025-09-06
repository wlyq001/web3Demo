package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PostApi struct{}

var PostApiApp = new(PostApi)

func (p *PostApi) CreatePost(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	post := &Post{}
	err := c.ShouldBindBodyWith(&post, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserId = uint(userId.(float64))
	DB.Create(post)
	c.JSON(http.StatusOK, "文章创建成功")
}

func (p *PostApi) GetPostList(c *gin.Context) {
	postList := []Post{}
	DB.Model(&Post{}).Select("id,title").Where("1=1").Order("created_at desc").Find(&postList)
	c.JSON(http.StatusOK, postList)
}

func (p *PostApi) GetPost(c *gin.Context) {
	post := &Post{}
	c.ShouldBindWith(post, binding.Query)
	DB.Model(&Post{}).Preload("User").First(post)
	c.JSON(http.StatusOK, post)
}

func (p *PostApi) EditPost(c *gin.Context) {
	postReq := &Post{}
	c.ShouldBindJSON(postReq)
	userId, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	postDB := &Post{}
	DB.Model(&Post{}).Where("id = ?", postReq.ID).First(postDB)

	if postDB.UserId != uint(userId.(float64)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	DB.Model(postReq).Select("title", "content").Updates(postReq)
	c.JSON(http.StatusOK, postReq)
}

func (p *PostApi) DeletePost(c *gin.Context) {
	postReq := &Post{}
	c.ShouldBindJSON(postReq)
	userId, exists := c.Get("userID")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	postDB := &Post{}
	DB.Model(&Post{}).Where("id = ?", postReq.ID).First(postDB)

	if postDB.UserId != uint(userId.(float64)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	DB.Delete(postReq)
	c.JSON(http.StatusOK, "删除成功")

}
