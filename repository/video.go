package repository

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserID        uint   `gorm:"not null"`
	Title         string `gorm:"not null"`
	FavoriteCount int    `gorm:"not null;default:0"`
	CommentCount  int    `gorm:"not null;default:0"`
	PublishTime   int64  `gorm:"not null"` //发布时间,视频流接口需要此属性
	//视频 封面
}

func VideoInsert(userID uint, title string) (int, error) {
	video := Video{
		UserID:      userID,
		Title:       title,
		PublishTime: time.Now().Unix()} //视频的发布时间
	res := db.Create(&video)
	if res.Error != nil {
		return 0, res.Error
	}
	var user User
	res = db.Model(&user).Where("id = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1))
	if res.Error != nil {
		return 0, res.Error
	}
	return int(video.ID), nil
}

// 查询发布视频
func VideoSelectUid(userID uint) (*[]Video, error) {
	var videos []Video
	res := db.Where("user_id = ?", userID).Find(&videos)
	if res.Error != nil {
		return nil, res.Error
	}
	return &videos, nil
}

// 查询最近30条视频
func VideoSelectAll(publishTime int64) (*[]Video, error) {
	var videos []Video
	res := db.Where("publish_time < ?", publishTime).Limit(30).Order("publish_time desc").Find(&videos)
	if res.Error != nil {
		return nil, res.Error
	}
	return &videos, nil
}
