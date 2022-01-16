package controllers

import "github.com/gofiber/fiber/v2"

func GetMenus(c *fiber.Ctx) error {
	return c.SendString("All Menus")
}

func GetMenu(c *fiber.Ctx) error {
	return c.SendString("Get Single Menu")
}

func CreateMenu(c *fiber.Ctx) error {
	return c.SendString("Create Menu")
}

func UpdateMenu(c *fiber.Ctx) error {
	return c.SendString("Update Menu")
}
