package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

type CommentApiRouter struct{}

func (r *CommentApiRouter) InitCommentApiRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("comment").Use(middleware.JWTAuth())
	normalRouter := Router.Group("comment")

	{
		jwtRouter.POST("createComment", commentApi.CreateComment)
	}
	{
		normalRouter.GET("commentList", commentApi.CommentList)
	}
}
