package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	UserID   uint `gorm:"not null"`
	ToUserID uint `gorm:"not null"`
}

func RelationSelectUidTouid(tokenUid int, uid int) (bool, error) {
	var relation Relation
	var isFollow bool
	res := db.Where("user_id = ? AND to_user_id = ?", tokenUid, uid).Take(&relation)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		isFollow = false
	} else if res.Error != nil {
		return false, res.Error
	} else {
		isFollow = true
	}
	return isFollow, nil
}

func AddFollow(uid int, to_uid int) (bool, error) {
	var relation Relation
	res := db.Where("user_id = ? AND to_user_id = ?", uid, to_uid).Take(&relation)
	if res.Error == nil { // 已关注
		return false, errors.New("不能重复关注")
	}
	res = db.Create(&Relation{UserID: uint(uid), ToUserID: uint(to_uid)})
	if res.Error != nil { // 关注失败
		return false, errors.New("关注失败")
	}
	user := User{}
	res = db.Model(&user).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count + ?", 1))
	if res.Error != nil {
		return false, res.Error
	}
	user = User{}
	res = db.Model(&user).Where("id = ?", to_uid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func CancelFollow(uid int, to_uid int) (bool, error) {
	var relation Relation
	res := db.Where("user_id = ? AND to_user_id = ?", uid, to_uid).Take(&relation)
	if res.Error != nil { // 未关注
		return false, errors.New("还未关注该用户")
	}
	res = db.Unscoped().Where("user_id = ? AND to_user_id = ?", uid, to_uid).Delete(&relation) // 硬删除
	if res.Error != nil {                                                                      // 取消关注失败
		return false, errors.New("取消关注失败")
	}
	user := User{}
	res = db.Model(&user).Where("id = ?", uid).Update("follow_count", gorm.Expr("follow_count - ?", 1))
	if res.Error != nil {
		return false, res.Error
	}
	user = User{}
	res = db.Model(&user).Where("id = ?", to_uid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

// 查询uid的关注列表
func RelationSelectToUidByUid(uid int) (*[]Relation, error) {
	var List []Relation
	res := db.Select("to_user_id").Where("user_id = ?", uid).Find(&List)
	if res.Error != nil {
		return nil, res.Error
	}
	return &List, nil
}

// 查询to_uid的粉丝(被关注)列表
func RelationSelectUidByToUid(toUid int) (*[]Relation, error) {
	var List []Relation
	res := db.Select("user_id").Where("to_user_id = ?", toUid).Find(&List)
	if res.Error != nil {
		return nil, res.Error
	}
	return &List, nil
}
