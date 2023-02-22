package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	VideoID uint `gorm:"not null"`
}

func FavoriteSelectUidVid(uid int, vid int) (bool, error) {
	var favorite Favorite
	var isFavorite bool
	res := db.Where("user_id = ? AND video_id = ?", uid, vid).Take(&favorite)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		isFavorite = false
	} else if res.Error != nil {
		return false, res.Error
	} else {
		isFavorite = true
	}
	return isFavorite, nil
}

func FavoriteSelectUid(uid int) (*[]Video, error) {
	var favorites []Favorite
	res := db.Select("video_id").Where("user_id = ?", uid).Find(&favorites)
	if res.Error != nil {
		return nil, res.Error
	}
	videos := make([]Video, 0, len(favorites))
	for _, favorite := range favorites {
		var video Video
		res = db.Where("id = ?", favorite.VideoID).Take(&video)
		if res.Error != nil {
			return nil, res.Error
		}
		videos = append(videos, video)
	}
	return &videos, nil
}

func UpdateFavorite(uid uint, vid uint) {
	isFavorite, _ := FavoriteSelectUidVid(int(uid), int(vid))
	var favorite Favorite
	var video Video
	var user User
	// 已经点赞，再点取消点赞
	if isFavorite {
		db.Model(&favorite).Where("user_id = ? AND video_id = ?", uid, vid).Update("deleted_at", time.Now())
		db.Model(&video).Where("id = ?", vid).Update("favorite_count", gorm.Expr("favorite_count - 1"))
		db.Model(&user).Where("id = ?", uid).Update("favorite_count", gorm.Expr("favorite_count - 1"))
		video = Video{}
		db.Where("id = ?", vid).Take(&video)
		user = User{}
		db.Model(&user).Where("id = ?", video.UserID).Update("total_favorited", gorm.Expr("total_favorited - 1"))
		return
	} else {
		db.Model(&favorite).Create(&Favorite{UserID: uid, VideoID: vid})
		db.Model(&video).Where("id = ?", vid).Update("favorite_count", gorm.Expr("favorite_count + 1"))
		db.Model(&user).Where("id = ?", uid).Update("favorite_count", gorm.Expr("favorite_count + 1"))
		video = Video{}
		db.Where("id = ?", vid).Take(&video)
		user = User{}
		db.Model(&user).Where("id = ?", video.UserID).Update("total_favorited", gorm.Expr("total_favorited + 1"))
	}
}
