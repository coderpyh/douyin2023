package service

import (
	"douyin2023/repository"
	"errors"
	"strconv"
	"time"
)

type FeedRequest struct {
	LatestTime int64  `query:"latest_time"`
	Token      string `query:"token"`
	TokenUid   int
}

type VideoResponse struct {
	VideoID       int          `json:"id"`
	Author        UserResponse `json:"author"`
	PlayURL       string       `json:"play_url"`
	CoverURL      string       `json:"cover_url"`
	FavoriteCount int          `json:"favorite_count"`
	CommentCount  int          `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

var LocalAddr string

// Feed 获取视频列表
func Feed(request *FeedRequest) (int64, *[]VideoResponse, error) {
	err := feedCheckRequest(request)
	if err != nil {
		return 0, nil, err
	}
	nextTime, response, err := feedPackResponse(request)
	if err != nil {
		return 0, nil, err
	}
	return nextTime, response, nil
}

func feedCheckRequest(request *FeedRequest) error {
	if request.LatestTime == 0 {
		request.LatestTime = time.Now().Unix()
	}
	if len(request.Token) > 0 {
		tokenUid, err := tokenToUid(request.Token)
		if err != nil {
			return errors.New("invalid param")
		}
		request.TokenUid = tokenUid
	}
	return nil
}

func feedPackResponse(request *FeedRequest) (int64, *[]VideoResponse, error) {
	videos, err := repository.VideoSelectAll(request.LatestTime)
	if err != nil {
		return 0, nil, err
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
			return 0, nil, err
		}
		response.Author = *userResponse
		response.PlayURL = "http://" + LocalAddr + "/douyin/data/" + strconv.Itoa(int(video.ID)) + ".mp4"
		response.CoverURL = "http://" + LocalAddr + "/douyin/data/0.jpg"
		response.FavoriteCount = video.FavoriteCount
		response.CommentCount = video.CommentCount
		isFavorite, err = repository.FavoriteSelectUidVid(int(video.UserID), int(video.ID))
		if err != nil {
			return 0, nil, err
		}
		response.IsFavorite = isFavorite
		response.Title = video.Title
		responses = append(responses, response)
	}
	var nextTime int64
	if len(*videos) != 0 {
		nextTime = (*videos)[len(*videos)-1].PublishTime
	}
	return nextTime, &responses, nil
}
