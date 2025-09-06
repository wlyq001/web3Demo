package router

import api "blog/service"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	BaseApiRouter
	PostApiRouter
	CommentApiRouter
}

var (
	baseApi    = api.BaseApiApp
	postApi    = api.PostApiApp
	commentApi = api.CommentApiApp
)
