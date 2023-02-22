package service

import (
	"crypto/sha256"
	"douyin2023/repository"
	"encoding/hex"
	"errors"
	"strconv"
	"strings"
	"sync"
)

type UserRequest struct {
	UserID   int    `query:"user_id"`
	Token    string `query:"token"`
	TokenUid int
}

type UserResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int    `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoriteCount   int    `json:"favorite_count"`
}

func User(request *UserRequest) (*UserResponse, error) {
	err := userCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := userPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func userCheckRequest(request *UserRequest) error {
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

// goroutine协程 分别到两个数据表中查询用户信息和是否关注
func userPackResponse(request *UserRequest) (*UserResponse, error) {
	var wg sync.WaitGroup
	var user *repository.User
	var isFollow bool
	var err1, err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		user, err1 = repository.UserSelectUid(request.UserID)
	}()
	go func() {
		defer wg.Done()
		isFollow, err2 = repository.RelationSelectUidTouid(request.TokenUid, request.UserID)
	}()
	wg.Wait()
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	var response UserResponse
	response.ID = request.UserID
	response.Name = user.Name
	response.FollowCount = user.FollowCount
	response.FollowerCount = user.FollowerCount
	response.IsFollow = isFollow
	response.Avatar = "http://" + LocalAddr + "/douyin/data/0.jpg"
	response.BackgroundImage = "http://" + LocalAddr + "/douyin/data/0.jpg"
	response.Signature = "个人简介"
	response.TotalFavorited = user.TotalFavorited
	response.WorkCount = user.WorkCount
	response.FavoriteCount = user.FavoriteCount
	return &response, nil
}

// token=uid_[uid加盐hash值]
func uidToToken(uid int) string {
	uidHashByte := sha256.Sum256([]byte("a" + strconv.Itoa(uid)))
	return strconv.Itoa(uid) + "_" + hex.EncodeToString(uidHashByte[:])
}

// 检查token是否合法并取出uid，没有在数据库中查找uid是否存在
func tokenToUid(token string) (int, error) {
	uidString, _, _ := strings.Cut(token, "_")
	uid, err := strconv.Atoi(uidString)
	if err != nil || uid <= 0 {
		return 0, errors.New("invalid param")
	}
	if uidToToken(uid) != token {
		return 0, errors.New("invalid param")
	}
	return uid, nil
}
