package repository

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	VideoID uint   `gorm:"not null"`
	Content string `gorm:"not null"`
}

func CommentSelectVid(vid int) (*[]Comment, error) {
	var comments []Comment
	res := db.Where("video_id = ?", vid).Order("created_at desc").Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return &comments, nil
}

func AddComment(uid, vid uint, content string) int {
	comment := Comment{
		UserID:  uid,
		VideoID: vid,
		Content: content,
	}
	db.Create(&comment)
	var video Video
	db.Model(&video).Where("id = ?", vid).Update("comment_count", gorm.Expr("comment_count + 1"))
	return int(comment.ID)
}

func CommentDelete(cid int, vid int) {
	var comment Comment
	db.Where("id = ?", cid).Delete(&comment)
	var video Video
	db.Model(&video).Where("id = ?", vid).Update("comment_count", gorm.Expr("comment_count - 1"))
}
