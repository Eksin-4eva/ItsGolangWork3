package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:UserId"` //连接user和userid
	UserId    uint   `gorm:"not null"`
	Title     string `gorm:"type:varchar(255);not null"`
	Content   string `gorm:"type:longtext"`
	Status    int    `gorm:"default:0"`
	View      uint64 `gorm:"default:0"`
	StartTime int64
	EndTime   int64
}
