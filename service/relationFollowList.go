package service

import (
	"douyin2023/repository"
)

type RelationFollowListRequest struct {
	UserID int    `query:"user_id"`
	Token  string `query:"token"`
}

func RelationFollowList(request *RelationFollowListRequest) (*[]UserResponse, error) {

	relationList, err := repository.RelationSelectToUidByUid(request.UserID)
	if err != nil {
		return nil, err
	}

	followList := make([]UserResponse, len(*relationList))
	for i, relation := range *relationList {
		userRequest := UserRequest{UserID: int(relation.ToUserID), Token: request.Token}
		userResponsePtr, err := User(&userRequest)
		if err != nil {
			return nil, err
		}
		followList[i] = *userResponsePtr
	}

	return &followList, nil
}
