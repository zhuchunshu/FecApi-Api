/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/Middleware"
	v1 "github.com/zhuchunshu/FecApi-Api/routes/api/v1"
)

func InitRouters(app *fiber.App) {
	api := app.Group("/api",Middleware.Api) // /api
	v1api:=api.Group("/v1",Middleware.ApiV1)
	v1.RegisterApiV1Routes(v1api)
}