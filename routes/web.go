package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/FecApi-Api/helpers/response"
	"github.com/zhuchunshu/FecApi-Api/helpers/types"
)

type Map map[string]interface{}

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

	route.Get("/test", func(ctx *fiber.Ctx) error {
		data := config.GetDatabaseConfig()
		MysqlConnect := helpers.Json_decode(data, "connect.mysql")
		ctx.Type("json", "utf-8")

		var datas types.Map = helpers.JsonToMap(helpers.Json_decode(data, "mysql."+MysqlConnect))

		return ctx.JSON(datas)
	})

	route.Get("/test/1", func(ctx *fiber.Ctx) error {
		return ctx.JSON(response.JsonApi(200,"success",types.Map{"message": "成功"})["result"])
	})

}
