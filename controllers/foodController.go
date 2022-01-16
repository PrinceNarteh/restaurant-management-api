package controllers

import "github.com/gofiber/fiber/v2"

func GetAllFoods(c *fiber.Ctx) error {
	return c.SendString("All Foods")
}
