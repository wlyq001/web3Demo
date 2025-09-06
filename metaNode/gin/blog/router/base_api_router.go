package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

type BaseApiRouter struct{}

func (e *BaseApiRouter) InitBaseApiRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("baseApi").Use(middleware.JWTAuth()).Use(middleware.Logger())
	normalRouter := Router.Group("baseApi").Use(middleware.Logger())
	{
		normalRouter.POST("register", baseApi.Register)
		normalRouter.POST("login", baseApi.Login)

	}
	{
		jwtRouter.GET("jwtTest", baseApi.JwtTest)
	}
}
