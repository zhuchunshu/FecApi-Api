package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/FecApi-Api/routes"
)

// 服务配置
var ServerConfig string = config.GetServerConfig()

func main() {
	app := fiber.New()

	// 注册路由
	routes.InitRouters(app)

	// 设置端口
	Serverport := helpers.Json_decode(ServerConfig, "port")
	app.Listen(":" + Serverport)
}
