package controllers

import "github.com/gofiber/fiber/v2"

func GetFoods(c *fiber.Ctx) error {
	return c.SendString("All Foods")
}

func GetFood(c *fiber.Ctx) error {
	return c.SendString("Get Single Food")
}

func CreateFood(c *fiber.Ctx) error {
	return c.SendString("Create Food")
}

func UpdateFood(c *fiber.Ctx) error {
	return c.SendString("Update Food")
}
