package service

import (
	"douyin2023/repository"
	"errors"
	"strconv"
)

type PublishListRequest struct {
	Token    string `query:"token"`
	UserID   int    `query:"user_id"`
	TokenUid int
}

// PublishList 获取用户发布的视频列表
func PublishList(request *PublishListRequest) (*[]VideoResponse, error) {
	err := publishListCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := publishListPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func publishListCheckRequest(request *PublishListRequest) error {
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

func publishListPackResponse(request *PublishListRequest) (*[]VideoResponse, error) {
	videos, err := repository.VideoSelectUid(uint(request.UserID))
	if err != nil {
		return nil, err
	}
	responses := make([]VideoResponse, 0, len(*videos))
	var response VideoResponse
	var userRequest UserRequest
	userRequest.TokenUid = request.TokenUid
	var isFavorite bool
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
		isFavorite, err = repository.FavoriteSelectUidVid(int(video.UserID), int(video.ID))
		if err != nil {
			return nil, err
		}
		response.IsFavorite = isFavorite
		response.Title = video.Title
		responses = append(responses, response)
	}
	return &responses, nil
}
