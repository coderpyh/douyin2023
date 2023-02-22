package repository

import (
	"gorm.io/gorm"
)

// 保存相关数量信息，避免每次查询重新统计
type User struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Password       string `gorm:"not null"`
	FollowCount    int    `gorm:"not null;default:0"`
	FollowerCount  int    `gorm:"not null;default:0"`
	TotalFavorited int    `gorm:"not null;default:0"`
	WorkCount      int    `gorm:"not null;default:0"`
	FavoriteCount  int    `gorm:"not null;default:0"`
	//头像 背景图 简介
}

func UserInsert(name string, password string) (int, error) {
	user := User{
		Name:     name,
		Password: password}
	res := db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(user.ID), nil
}

func UserSelectUid(uid int) (*User, error) {
	var user User
	res := db.
		Select("name", "follow_count", "follower_count", "total_favorited", "work_count", "favorite_count").
		Take(&user, uid)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func UserSelectNamePassword(name string, password string) (int, error) {
	var user User
	res := db.Select("id").Where("name = ? AND password = ?", name, password).Take(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return int(user.ID), nil
}
