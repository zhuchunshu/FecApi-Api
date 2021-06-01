/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "time"

type AdminSwitch struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name   string `json:"name"`
	Status int
}

func (AdminSwitch) TableName() string {
	return "admin_switch"
}