package main

import (
	"douyin2023/repository"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	repository.Init()
	h := server.Default()
	initRouter(h)
	h.Spin()
}
