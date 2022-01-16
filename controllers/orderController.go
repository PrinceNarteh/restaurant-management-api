package controllers

import "github.com/gofiber/fiber/v2"

func GetOrders(c *fiber.Ctx) error {
	return c.SendString("All Orders")
}

func GetOrder(c *fiber.Ctx) error {
	return c.SendString("Get Single Order")
}

func CreateOrder(c *fiber.Ctx) error {
	return c.SendString("Create Order")
}

func UpdateOrder(c *fiber.Ctx) error {
	return c.SendString("Update Order")
}
