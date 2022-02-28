package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `gorm:"index;not null"` //加索引加快查询速度
	Status uint `gorm:"default:0"` // 0 未完成 1 已完成
	Content string `gorm:"type longtext"`
	StartTime int64
	EndTime int64
	User User
	UserID int
}