package service

import (
	"douyin2023/repository"
)

type RelationFollowerListRequest struct {
	UserID int    `query:"user_id"`
	Token  string `query:"token"`
}

func RelationFollowerList(request *RelationFollowerListRequest) (*[]UserResponse, error) {

	relationList, err := repository.RelationSelectUidByToUid(request.UserID)
	if err != nil {
		return nil, err
	}

	followerList := make([]UserResponse, len(*relationList))
	for i, relation := range *relationList {
		userRequest := UserRequest{UserID: int(relation.UserID), Token: request.Token}
		userResponsePtr, err := User(&userRequest)
		if err != nil {
			return nil, err
		}
		followerList[i] = *userResponsePtr
	}

	return &followerList, nil
}
