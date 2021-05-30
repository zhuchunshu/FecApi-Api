/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Database

import "github.com/zhuchunshu/FecApi-Api/app/server/database"

type Users struct {
	database.TypeModel
	Name  string `json:"name"`
}