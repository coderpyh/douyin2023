package controller

import (
	"context"
	"douyin2023/service"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func UserLogin(c context.Context, ctx *app.RequestContext) {
	var request service.UserLoginRequest
	err := ctx.BindAndValidate(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"user_id":     0,
			"token":       "0"})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}
	response, err := service.UserLogin(&request)
	if err != nil {
		ctx.JSON(consts.StatusOK, utils.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"user_id":     0,
			"token":       "0"})
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}
	ctx.JSON(consts.StatusOK, utils.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_id":     response.UserID,
		"token":       response.Token})
	fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "success", response)
}
