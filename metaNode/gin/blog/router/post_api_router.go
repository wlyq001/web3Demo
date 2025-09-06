package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

type PostApiRouter struct{}

func (p *PostApiRouter) InitPostApiRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("post").Use(middleware.JWTAuth())
	normalRouter := Router.Group("post")

	{
		jwtRouter.POST("createPost", postApi.CreatePost)
		jwtRouter.POST("editPost", postApi.EditPost)
		jwtRouter.POST("deletePost", postApi.DeletePost)
	}
	{
		normalRouter.GET("getPostList", postApi.GetPostList)
		normalRouter.GET("getPost", postApi.GetPost)
	}
}
