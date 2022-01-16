package controllers

import "github.com/gofiber/fiber/v2"

func GetTables(c *fiber.Ctx) error {
	return c.SendString("All Tables")
}

func GetTable(c *fiber.Ctx) error {
	return c.SendString("Get Single Table")
}

func CreateTable(c *fiber.Ctx) error {
	return c.SendString("Create Table")
}

func UpdateTable(c *fiber.Ctx) error {
	return c.SendString("Update Table")
}
