/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "github.com/zhuchunshu/FecApi-Api/app/server/database"

type PersonalAccessToken struct {
	database.TypeModel
	UserId    uint `gorm:"column:tokenable_id"`
	Name      string
	Token     string
	Abilities string
	User      User `gorm:"foreignKey:UserId"`
}


func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens"
}