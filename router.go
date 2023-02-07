package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func middleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(ctx.FullPath(), string(ctx.Method()), ctx.ClientIP(), "start")
		ctx.Next(c)
		fmt.Println(ctx.FullPath(), string(ctx.Method()), ctx.ClientIP(), "end")
	}
}

func initRouter(h *server.Hertz) {
	h.Use(middleware())
	douyin := h.Group("/douyin")
	douyin.GET("/feed/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"feed": "feed"})
	})
	user := douyin.Group("/user")
	user.POST("/register/", func(c context.Context, ctx *app.RequestContext) {
		fmt.Println("abc")
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "string",
			"user_id":     0,
			"token":       "string"})
	})
	user.POST("/login/", controller.userRegister)
	controller.userRegister()
	user.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 0,
			"status_msg":  "string",
			"user": utils.H{
				"id":             0,
				"name":           "string",
				"follow_count":   0,
				"follower_count": 0,
				"is_follow":      true}})
	})
}
