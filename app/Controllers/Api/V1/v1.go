/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package V1

import (
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/zhuchunshu/FecApi-Api/app"
	v11 "github.com/zhuchunshu/FecApi-Api/app/Jobs/V1"
	"github.com/zhuchunshu/FecApi-Api/app/Models/Database"
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
	"github.com/zhuchunshu/FecApi-Api/helpers/String"
	database2 "github.com/zhuchunshu/FecApi-Api/helpers/database"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
)

func QQHOOK(ctx *fiber.Ctx) error {
	token := database2.GetOptions("api-v1-qqhook-token")
	urls := database2.GetOptions("api-v1-qqhook-url")
	if len(token) != 0 && len(urls) != 0 {
		qq := ctx.Query("qq")
		content := ctx.Query("content")
		title := ctx.Query("title", "新通知!")
		if len(qq) == 0 {
			return ctx.JSON(response.StringApi(200, true, "缺少参数: qq"))
		}
		if len(content) == 0 {
			return ctx.JSON(response.StringApi(200, true, "缺少参数: content (通知内容)"))
		}
		db := database.DBConn
		hash, _ := gonanoid.New()
		apinotice := Database.ApiNotice{
			Title:   title,
			Content: content,
			Hash:    hash,
			UserId:app.ApiData.UserId,
		}
		db.Create(&apinotice)
		t:=String.Base64Encode([]byte(title+"\n\n\n查看详情:暂无地址"))

		go v11.QQHook(urls,qq,t,token)
		return ctx.JSON(response.StringApi(200,true,"通知信息已发送!"))
	} else {
		return ctx.JSON(response.StringApi(404, false, "此接口服务端配置不完整"))
	}
}
