package controller

import (
	"context"
	"douyin2023/service"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func MessageChat(c context.Context, ctx *app.RequestContext) {
	var request service.MessageChatRequest
	err := ctx.BindAndValidate(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code":  1,
			"status_msg":   err.Error(),
			"message_list": nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}

	response, err := service.MessageChat(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code":  1,
			"status_msg":   err.Error(),
			"message_list": nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}
	ctx.JSON(consts.StatusOK, utils.H{
		"status_code":  0,
		"status_msg":   nil,
		"message_list": response})
	fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "success", response)
}
