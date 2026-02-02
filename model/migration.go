package model

import (
	"fmt"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Task{})
	if err != nil {
		fmt.Println("Migration failed:" + err.Error())
	}
}
