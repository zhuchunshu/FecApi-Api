/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import (
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
	"time"
)

type Users struct {
	database.TypeModel
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"profile_photo_path"`
}

type User struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"profile_photo_path"`
}


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
func (User) TableName() string {
	return "users"
}
