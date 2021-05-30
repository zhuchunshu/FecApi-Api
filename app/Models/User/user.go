/*
 * MIT License
 *
 * Copyright (c) 2021 朱纯树
 */

package User

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/FecApi-Api/app/Models/Database"
	"github.com/zhuchunshu/FecApi-Api/app/server/database"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var users []Database.Users
	db.Find(&users)
	return c.JSON(users)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var User Database.Users
	db.Find(&User, id)
	return c.JSON(User)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Database.Users)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

