package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/routes/api"
)

func InitRouters(app *fiber.App) {
	RegisterWebRoutes(app)
	api.InitRouters(app)
}
