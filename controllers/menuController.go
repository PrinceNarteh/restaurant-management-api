package controllers

import (
	"context"
	"github.com/PrinceNarteh/restaurant-management-api/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func GetMenus(c *fiber.Ctx) error {
	result, err := database.MenuCollection.Find(ctx, bson.D{})
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get menus"})
	}
	var menus bson.M
	if err = result.All(ctx, &menus); err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": menus})
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
