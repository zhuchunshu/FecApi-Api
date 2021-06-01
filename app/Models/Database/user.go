/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "time"

type User struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"profile_photo_path"`
}

func (User) TableName() string {
	return "users"
}