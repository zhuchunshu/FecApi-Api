/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "time"

type AdminSetting struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name   string `json:"name"`
	Value string
}

func (AdminSetting) TableName() string {
	return "admin_setting"
}