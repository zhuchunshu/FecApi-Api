/*
 * MIT License
 *	数据库辅助函数
 * Copyright (c) 2021 朱纯树
 */

package database

import (
	"github.com/zhuchunshu/FecApi-Api/app/Models/Database"
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
)

func GetSwitch(name string)bool{
	db := database.DBConn
	var count int64
	var AdminSwitch Database.AdminSwitch
	db.Model(&AdminSwitch).Where("name = ?", name).Count(&count)
	if count>=1 {
		return true
	}else {
		return false
	}
}

// 获取站点设置项
func GetOptions(name string)string{
	db:=database.DBConn
	var AdminSetting Database.AdminSetting
	db.Where("name = ?",name).First(&AdminSetting)
	return AdminSetting.Value
}