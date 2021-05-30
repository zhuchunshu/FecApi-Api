/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/Controllers/Api/V1"
)

func  RegisterApiV1Routes(route fiber.Router)  {
	route.Post("/test", V1.TestPage)
}