package service

import (
	"douyin2023/repository"
	"errors"
)

type FavoriteRequest struct {
	Token      string `query:"token"`
	VideoID    uint   `query:"video_id"`
	ActionType int    `query:"action_type"`
}

func AddFavorite(req *FavoriteRequest) error {
	if req.VideoID == 0 {
		return errors.New("invalid param")
	}
	uid, err := tokenToUid(req.Token)
	if err != nil {
		return err
	}
	repository.UpdateFavorite(uint(uid), req.VideoID)
	return nil
}
