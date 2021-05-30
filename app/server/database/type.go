/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package database

import "time"

type TypeModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}