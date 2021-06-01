/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "github.com/zhuchunshu/FecApi-Api/app/server/database"

type ApiNotice struct {
	database.TypeModel
	Title    string
	Content      string
	Hash     string
	UserId uint `gorm:"column:user_id"`
}


func (ApiNotice) TableName() string {
	return "api_notice"
}