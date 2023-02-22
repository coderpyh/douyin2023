package service

import (
	"douyin2023/repository"
	"errors"
)

type CommentListRequest struct {
	Token    string `query:"token"`
	VideoID  int    `query:"video_id"`
	TokenUid int
}

type CommentResponse struct {
	ID         int          `json:"id"`
	User       UserResponse `json:"user"`
	Content    string       `json:"content"`
	CreateDate string       `json:"create_date"`
}

func CommentList(request *CommentListRequest) (*[]CommentResponse, error) {
	err := commentListCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := commentListPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func commentListCheckRequest(request *CommentListRequest) error {
	if request.VideoID <= 0 {
		return errors.New("invalid param")
	}
	tokenUid, err := tokenToUid(request.Token)
	if err != nil {
		return errors.New("invalid param")
	}
	request.TokenUid = tokenUid
	return nil
}

func commentListPackResponse(request *CommentListRequest) (*[]CommentResponse, error) {
	comments, err := repository.CommentSelectVid(request.VideoID)
	if err != nil {
		return nil, err
	}
	responses := make([]CommentResponse, 0, len(*comments))
	var response CommentResponse
	var userRequest UserRequest
	userRequest.TokenUid = request.TokenUid
	for _, comment := range *comments {
		response.ID = int(comment.ID)
		userRequest.UserID = int(comment.UserID)
		userResponse, err := userPackResponse(&userRequest)
		if err != nil {
			return nil, err
		}
		response.User = *userResponse
		response.Content = comment.Content
		response.CreateDate = comment.CreatedAt.Format("01-02")
		responses = append(responses, response)
	}
	return &responses, nil
}
