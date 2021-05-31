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
	"gorm.io/gorm/schema"
)

func initDatabase() {
	var err error
	var GetDatabaseConfig string = config.GetDatabaseConfig()
	var MysqlConnect string = helpers.JsonDecode(GetDatabaseConfig, "connect.mysql")
	var datas string = helpers.JsonDecode(GetDatabaseConfig, "mysql."+MysqlConnect)
	var db string = helpers.JsonDecode(datas, "database")
	var user string = helpers.JsonDecode(datas, "username")
	var host string = helpers.JsonDecode(datas, "host")
	var port string = helpers.JsonDecode(datas, "port")
	var pwd string = helpers.JsonDecode(datas, "password")
	// 数据库连接配置
	dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建数据库连接
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, //禁用默认事务
		PrepareStmt: true,//缓存预编译语句

		NamingStrategy: schema.NamingStrategy{
			SingularTable:true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	fmt.Println("Database Migrated")
}

