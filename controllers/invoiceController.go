package controllers

import "github.com/gofiber/fiber/v2"

func GetInvoices(c *fiber.Ctx) error {
	return c.SendString("All Invoices")
}

func GetInvoice(c *fiber.Ctx) error {
	return c.SendString("Get Single Invoice")
}

func CreateInvoice(c *fiber.Ctx) error {
	return c.SendString("Create Invoice")
}

func UpdateInvoice(c *fiber.Ctx) error {
	return c.SendString("Update Invoice")
}
