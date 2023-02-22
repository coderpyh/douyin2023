package service

import (
	"douyin2023/repository"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

type PublishActionRequest struct {
	fileHeader *multipart.FileHeader `form:"data"`
	Token      string                `form:"token"`
	Title      string                `form:"title"`
}

func PublishAction(request *PublishActionRequest) error {
	fmt.Println("token =", request.Token, "title =", request.Title)

	open, err := request.fileHeader.Open()
	if err != nil {
		return err
	}
	// 读取文件到字节数组
	fileRaw, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}

	uid, _, _ := strings.Cut(request.Token, "_")

	// 将文件名写入到数据库
	userID, _ := strconv.Atoi(uid)
	videoID, err := repository.VideoInsert(uint(userID), request.Title)
	if err != nil {
		return err
	}

	// 将字节数组写入到文件
	// 生成文件名（vid）
	filename := "data/" + strconv.Itoa(videoID) + ".mp4"
	err = os.WriteFile(filename, fileRaw, 0666)
	if err != nil {
		return err
	}

	fmt.Println("videoID =", videoID)
	return nil
}
