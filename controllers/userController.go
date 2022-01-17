package controllers

import "github.com/gofiber/fiber/v2"

func GetUsers(ctx *fiber.Ctx) error {

	return ctx.SendString("Get Users")
}

func GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Get User")
}

func Register(ctx *fiber.Ctx) error {
	return ctx.SendString("Register User")
}

func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("Login User")
}
