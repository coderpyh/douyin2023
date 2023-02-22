package service

import (
	"douyin2023/repository"
	"strconv"
	"strings"
)

type MessageChatRequest struct {
	Token      string `query:"token"`
	ToUserID   int    `query:"to_user_id"`
	PreMsgTime int64  `query:"pre_msg_time"`
}

type MessageChatResponse struct {
	ID         uint   `json:"id"`
	ToUserID   uint   `json:"to_user_id"`
	FromUserID uint   `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func MessageChat(request *MessageChatRequest) (*[]MessageChatResponse, error) {
	fromUidStr, _, _ := strings.Cut(request.Token, "_")
	fromUserID, _ := strconv.Atoi(fromUidStr)
	messageList, err := repository.SelectMessage(uint(fromUserID), uint(request.ToUserID), request.PreMsgTime)
	if err != nil {
		return nil, err
	}

	response := make([]MessageChatResponse, len(*messageList))
	for i, message := range *messageList {
		response[i].ID = message.ID
		response[i].ToUserID = message.ToUserID
		response[i].FromUserID = message.FromUserID
		response[i].Content = message.Content
		response[i].CreateTime = message.CreatedAt.Unix()
	}
	return &response, nil
}
