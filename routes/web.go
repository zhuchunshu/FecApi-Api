package routes

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(route *fiber.App) {

	// 静态页面
	route.All("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"code":   200,
			"status": "success",
			"data": fiber.Map{
				"message": "FecApi is ok",
			},
		})
	})

}
