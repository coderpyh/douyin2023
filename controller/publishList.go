package controller

import (
	"context"
	"douyin2023/service"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PublishList(c context.Context, ctx *app.RequestContext) {
	var request service.PublishListRequest
	err := ctx.BindAndValidate(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"video_list":  nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}

	response, err := service.PublishList(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"video_list":  nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}

	//返回发布视频列表
	ctx.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"video_list":  response})
	fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "success", response)
}
