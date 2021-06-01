/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package V1

import (
	"github.com/gofiber/fiber/v2"
	database2 "github.com/zhuchunshu/FecApi-Api/helpers/database"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
	"net/http"
	"net/url"
)

func TestPage(ctx *fiber.Ctx)error{
	token:=database2.GetOptions("api-v1-qqhook-token")
	urls:=database2.GetOptions("api-v1-qqhook-url")
	if len(token)!=0 && len(urls)!=0 {
		qq :=ctx.Query("qq")
		text:=ctx.Query("text")
		if len(qq)==0 {
			return ctx.JSON(response.StringApi(200, true, "缺少参数: qq"))
		}
		if len(text)==0 {
			return ctx.JSON(response.StringApi(200, true, "缺少参数: text"))
		}
		http.PostForm(urls+token+"/private/"+qq+"/"+text,url.Values{})
	}else {
		return ctx.JSON(response.StringApi(404,false,"此接口服务端配置不完整"))
	}
	return ctx.JSON(response.StringApi(200,true,"1"))
}
