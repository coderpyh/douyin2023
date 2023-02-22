package main

import (
	"context"
	"douyin2023/controller"
	"douyin2023/service"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func printRequest() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "start")
		service.LocalAddr = ctx.GetConn().LocalAddr().String()
		ctx.Next(c)
	}
}

func initRouter(h *server.Hertz) {
	h.Use(printRequest())
	douyin := h.Group("/douyin")
	douyin.GET("/data/:filename", controller.Data)

	//基础
	feed := douyin.Group("/feed")
	feed.GET("/", controller.Feed)
	user := douyin.Group("/user")
	user.POST("/register/", controller.UserRegister)
	user.POST("/login/", controller.UserLogin)
	user.GET("/", controller.User)
	publish := douyin.Group("/publish")
	publish.POST("/action/", controller.PublishAction)
	publish.GET("/list/", controller.PublishList)

	//互动
	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", controller.FavoriteAction)
	favorite.GET("/list/", controller.FavoriteList)
	comment := douyin.Group("/comment")
	comment.POST("/action/", controller.CommentAction)
	comment.GET("/list/", controller.CommentList)

	//社交
	relation := douyin.Group("/relation")
	relation.POST("/action/", controller.RelationAction)
	relation.GET("/follow/list/", controller.RelationFollowList)
	relation.GET("/follower/list/", controller.RelationFollowerList)
	relation.GET("/friend/list/", controller.RelationFriendList)
	message := douyin.Group("/message")
	message.POST("/action/", controller.MessageAction)
	message.GET("/chat/", controller.MessageChat)
}
