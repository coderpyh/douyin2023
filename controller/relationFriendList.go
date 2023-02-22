package controller

import (
	"context"
	"douyin2023/service"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RelationFriendList(c context.Context, ctx *app.RequestContext) {

	var request service.RelationFriendListRequest
	err := ctx.BindAndValidate(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"user_list":   nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}

	friendList, err := service.RelationFriendList(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"user_list":   nil})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}

	ctx.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_list":   friendList})
	fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "success", friendList)
}
