package controllers

import "github.com/gofiber/fiber/v2"

func GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("Get User")
}

func GetUsers(ctx *fiber.Ctx) error {
	return ctx.SendString("Get Users")
}

func Register(ctx *fiber.Ctx) error {
	return ctx.SendString("Register User")
}

func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("Login User")
}
