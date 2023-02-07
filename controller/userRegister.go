package controller

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func userRegister(c context.Context, ctx *app.RequestContext) []byte {
	a := userInfo{Name: "wang", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	b, _ := json.Marshal(a)
	return b
}
