package service

import (
	"douyin2023/repository"

	"github.com/fatih/set" // 求交集
)

type RelationFriendListRequest struct {
	UserID int    `query:"user_id"`
	Token  string `query:"token"`
}

type RelationFriendListResponse struct {
	UserResponse
	Message string `json:"message"`
	MsgType int    `json:"msgType"`
}

func RelationFriendList(request *RelationFriendListRequest) (*[]RelationFriendListResponse, error) {

	followRelationList, err := repository.RelationSelectToUidByUid(request.UserID) // 关注列表
	if err != nil {
		return nil, err
	}

	followerRelationList, err := repository.RelationSelectUidByToUid(request.UserID) // 粉丝列表
	if err != nil {
		return nil, err
	}

	followIDSet := set.New(set.ThreadSafe)
	for _, relation := range *followRelationList {
		followIDSet.Add(relation.ToUserID)
	}

	followerIDSet := set.New(set.ThreadSafe)
	for _, relation := range *followerRelationList {
		followerIDSet.Add(relation.UserID)
	}

	friendIDSet := set.Intersection(followIDSet, followerIDSet)

	friendList := make([]RelationFriendListResponse, friendIDSet.Size())
	for i, friendID := range friendIDSet.List() {
		userRequest := UserRequest{UserID: int(friendID.(uint)), Token: request.Token}
		userResponsePtr, err := User(&userRequest)
		if err != nil {
			return nil, err
		}
		friendList[i].UserResponse = *userResponsePtr
		// 获取最新的一条消息及类型
		friendList[i].Message, friendList[i].MsgType = repository.SelectLatestMessage(uint(request.UserID), friendID.(uint))
	}

	return &friendList, nil
}
