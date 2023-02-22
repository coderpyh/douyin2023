package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(
		mysql.Open("root:root@tcp(localhost:3306)/douyin2023?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{})
	printError("连接数据库失败", err)
	err = db.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Comment{}, &Relation{}, &Message{})
	printError("建表失败", err)
}

func printError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
