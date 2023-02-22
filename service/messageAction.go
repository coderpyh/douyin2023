package service

import (
	"douyin2023/repository"
	"strconv"
	"strings"
)

type MessageActionRequest struct {
	Token      string `query:"token"`
	ToUserID   int    `query:"to_user_id"`
	ActionType int    `query:"action_type"`
	Content    string `query:"content"`
}

func MessageAction(request *MessageActionRequest) (bool, error) {
	fromUidStr, _, _ := strings.Cut(request.Token, "_")
	fromUid, _ := strconv.Atoi(fromUidStr)
	toUid := request.ToUserID
	content := request.Content
	return repository.SendMessage(uint(fromUid), uint(toUid), content)
}
