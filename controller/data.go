package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

type DataRequest struct {
	Filename string `path:"filename"`
}

func Data(c context.Context, ctx *app.RequestContext) {
	var request DataRequest
	err := ctx.BindAndValidate(&request)
	if err != nil {
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", err)
		return
	}
	if strings.HasSuffix(request.Filename, ".mp4") {
		ctx.SetContentType("video/mp4")
	} else if strings.HasSuffix(request.Filename, ".jpg") {
		ctx.SetContentType("image/jpeg")
	} else {
		fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "fail", "invalid param")
		return
	}
	ctx.File("data/" + request.Filename)
	fmt.Println(string(ctx.Method()), string(ctx.Request.RequestURI()), ctx.ClientIP(), "success")
}
