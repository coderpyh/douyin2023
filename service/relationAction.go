package service

import (
	"douyin2023/repository"
	"errors"
	"strconv"
	"strings"
)

type RelationActionRequest struct {
	Token      string `query:"token"`
	ToUserID   int    `query:"to_user_id"`
	ActionType int    `query:"action_type"`
}

func RelationAction(request *RelationActionRequest) (bool, error) {
	uidStr, _, _ := strings.Cut(request.Token, "_")
	uid, _ := strconv.Atoi(uidStr)
	to_uid := request.ToUserID
	if request.ActionType == 1 { // 添加关注
		return repository.AddFollow(uid, to_uid)
	} else if request.ActionType == 2 { // 取消关注
		return repository.CancelFollow(uid, to_uid)
	}
	return false, errors.New("invalid param")
}
