package service

import (
	"douyin2023/repository"
	"errors"
	"strconv"
)

type FavoriteListRequest struct {
	UserID   int    `query:"user_id"`
	Token    string `query:"token"`
	TokenUid int
}

func FavoriteList(request *FavoriteListRequest) (*[]VideoResponse, error) {
	err := favoriteListCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := favoriteListPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func favoriteListCheckRequest(request *FavoriteListRequest) error {
	if request.UserID <= 0 {
		return errors.New("invalid param")
	}
	tokenUid, err := tokenToUid(request.Token)
	if err != nil {
		return errors.New("invalid param")
	}
	request.TokenUid = tokenUid
	return nil
}

func favoriteListPackResponse(request *FavoriteListRequest) (*[]VideoResponse, error) {
	videos, err := repository.FavoriteSelectUid(request.UserID)
	if err != nil {
		return nil, err
	}
	//切片预分配内存
	responses := make([]VideoResponse, 0, len(*videos))
	var response VideoResponse
	var userRequest UserRequest
	userRequest.TokenUid = request.TokenUid
	for _, video := range *videos {
		response.VideoID = int(video.ID)
		userRequest.UserID = int(video.UserID)
		userResponse, err := userPackResponse(&userRequest)
		if err != nil {
			return nil, err
		}
		response.Author = *userResponse
		response.PlayURL = "http://" + LocalAddr + "/douyin/data/" + strconv.Itoa(int(video.ID)) + ".mp4"
		response.CoverURL = "http://" + LocalAddr + "/douyin/data/0.jpg"
		response.FavoriteCount = video.FavoriteCount
		response.CommentCount = video.CommentCount
		response.IsFavorite = true
		response.Title = video.Title
		responses = append(responses, response)
	}
	return &responses, nil
}
