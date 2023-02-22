package repository

import (
	"sort"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromUserID uint   `gorm:"not null"`
	ToUserID   uint   `gorm:"not null"`
	Content    string `gorm:"not null"`
}

// 发送消息
func SendMessage(fromUid uint, toUid uint, content string) (bool, error) {
	message := Message{
		FromUserID: fromUid,
		ToUserID:   toUid,
		Content:    content,
	}
	res := db.Create(&message)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

// 获取最新的一条消息
func SelectLatestMessage(fromUid uint, toUid uint) (string, int) {
	var message1, message2 Message
	db.Where("from_user_id = ? AND to_user_id = ?", fromUid, toUid).Order("created_at desc").Take(&message1)
	db.Where("from_user_id = ? AND to_user_id = ?", toUid, fromUid).Order("created_at desc").Take(&message2)
	if message1.CreatedAt.After(message2.CreatedAt) {
		return message1.Content, 1
	} else {
		return message2.Content, 0
	}
}

func SelectMessage(fromUid uint, toUid uint, preMsgTime int64) (*[]Message, error) {
	var messageList1, messageList2 []Message
	res := db.Where("from_user_id = ? AND to_user_id = ? AND created_at > ?", fromUid, toUid, preMsgTime).Find(&messageList1)
	if res.Error != nil {
		return nil, res.Error
	}
	res = db.Where("from_user_id = ? AND to_user_id = ? AND created_at > ?", toUid, fromUid, preMsgTime).Find(&messageList2)
	if res.Error != nil {
		return nil, res.Error
	}
	// 发送的消息与接受的消息合并
	messageList := append(messageList1, messageList2...)
	// 按时间排序
	sort.Slice(messageList, func(i, j int) bool {
		return messageList[i].CreatedAt.Before(messageList[j].CreatedAt)
	})
	return &messageList, nil
}
