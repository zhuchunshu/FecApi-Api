package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/server/config"
	"github.com/zhuchunshu/FecApi-Api/helpers"
	"github.com/zhuchunshu/FecApi-Api/routes"
	"log"
)

// 服务配置
var ServerConfig string = config.GetServerConfig()

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "FecApi",
	})

	// 注册路由
	routes.InitRouters(app)

	// 设置端口
	Serverport := helpers.Json_decode(ServerConfig, "port")
	server := app.Listen(":" + Serverport)
	log.Fatal(server)
}
