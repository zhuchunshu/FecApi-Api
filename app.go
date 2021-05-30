/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package main

import (
	"fmt"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	var GetDatabaseConfig string = config.GetDatabaseConfig()
	var MysqlConnect string = helpers.Json_decode(GetDatabaseConfig, "connect.mysql")
	var datas string = helpers.Json_decode(GetDatabaseConfig, "mysql."+MysqlConnect)
	var db string = helpers.Json_decode(datas, "database")
	var user string = helpers.Json_decode(datas, "username")
	var host string = helpers.Json_decode(datas, "host")
	var port string = helpers.Json_decode(datas, "port")
	var pwd string = helpers.Json_decode(datas, "password")
	dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	fmt.Println("Database Migrated")
}
