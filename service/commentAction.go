package service

import (
	"douyin2023/repository"
	"errors"
	"time"
)

type CommentRequest struct {
	Token       string `query:"token"`
	VideoID     uint   `query:"video_id"`
	ActionType  int    `query:"action_type"`
	CommentText string `query:"comment_text"`
	CommentID   uint   `query:"comment_id"`
}

func AddComment(req *CommentRequest) (*CommentResponse, error) {
	if req.ActionType == 1 {
		if req.CommentText == "" {
			return nil, errors.New("评论内容为空")
		}
		uid, _ := tokenToUid(req.Token)
		cid := repository.AddComment(uint(uid), req.VideoID, req.CommentText)
		var response CommentResponse
		response.ID = cid
		var userRequest UserRequest
		userRequest.TokenUid = uid
		userRequest.UserID = uid
		userResponse, err := userPackResponse(&userRequest)
		if err != nil {
			return nil, err
		}
		response.User = *userResponse
		response.Content = req.CommentText
		response.CreateDate = time.Now().Format("01-02")
		return &response, nil
	} else {
		repository.CommentDelete(int(req.CommentID), int(req.VideoID))
		return nil, nil
	}
}
