/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app"
	"github.com/zhuchunshu/FecApi-Api/app/Models/Database"
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
)

func Api(ctx *fiber.Ctx)error{
	accessToken :=ctx.Query("access_token","1")
	ctx.Set("access_token", accessToken)
	db := database.DBConn
	var PersonalAccessToken Database.PersonalAccessToken
	db.Where("token = ?", accessToken).Preload("User").First(&PersonalAccessToken)
	var count int64
	db.Where("token = ?", accessToken).Model(PersonalAccessToken).Count(&count)
	if count>=1 {
		if PersonalAccessToken.User.ID<=0 {
			return ctx.JSON(response.StringApi(401,false,"无权限"))
		}else {
			app.ApiData=PersonalAccessToken
			return ctx.Next()
		}
	}else{
		return ctx.JSON(response.StringApi(401,false,"无权限"))
	}
	//return ctx.Next()
}

func ApiV1(ctx *fiber.Ctx)error{
	ctx.Set("FecApi-Version", "v1")
	return ctx.Next()
}