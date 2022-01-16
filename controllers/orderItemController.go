package controllers

import "github.com/gofiber/fiber/v2"

func GetOrderItems(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Order Items")
}

func GetOrderItem(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Single Order Item")

}
func CreateOrderItem(ctx *fiber.Ctx) error {
	return ctx.SendString("Create Order Item")

}
func UpdateOrderItem(ctx *fiber.Ctx) error {
	return ctx.SendString("Update Order Item")

}
func GetOrderItemByOrder(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Order Item By Order")

}
