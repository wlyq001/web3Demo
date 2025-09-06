package main

import (
	"blog/router"
	"blog/service"
)

func main() {
	// 初始化数据库、gin
	service.Init()
	// 初始化路由
	group := service.ENGINE.Group("/")
	router.RouterGroupApp.BaseApiRouter.InitBaseApiRouter(group)
	router.RouterGroupApp.PostApiRouter.InitPostApiRouter(group)
	router.RouterGroupApp.CommentApiRouter.InitCommentApiRouter(group)
	service.RunGin()
}
