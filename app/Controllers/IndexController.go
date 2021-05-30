/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
)

// route -> /
func Index(ctx *fiber.Ctx)error{
	return ctx.JSON(response.StringApi(200,true,"FecApi is Ok!"))
}

func Tests(ctx *fiber.Ctx)error{
	return ctx.JSON(response.StringApi(200,true,"FecApi is Ok!"))
}