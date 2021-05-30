/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package V1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
)

func TestPage(ctx *fiber.Ctx)error{
	return ctx.JSON(response.StringApi(200,true,"test"))
}
