package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/Controllers"
)

type Map map[string]interface{}

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(route *fiber.App) {

	// 静态页面
	route.All("/",Controllers.Index)


	route.Get("/test/test",Controllers.Tests)
}
