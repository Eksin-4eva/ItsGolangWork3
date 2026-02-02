package model

import "gorm.io/gorm"

type User struct {
	gorm.Model            // 包含ID, CreatedAt, UpdatedAt, DeletedAt
	UserName       string `gorm:"unique"`
	PasswordDigest string `gorm:"comment:加密后的密码"` //加密后密码
}
