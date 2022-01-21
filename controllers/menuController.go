package controllers

import (
	"context"
	db "github.com/PrinceNarteh/restaurant-management-api/database"
	"github.com/PrinceNarteh/restaurant-management-api/helpers"
	"github.com/PrinceNarteh/restaurant-management-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func GetMenus(c *fiber.Ctx) error {
	result, err := db.MenuCollection.Find(ctx, bson.D{})
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not get menus"})
	}
	var menus bson.M
	if err = result.All(ctx, &menus); err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(menus)
}

func GetMenu(c *fiber.Ctx) error {
	var menu models.Menu
	menuId := c.Params("menuId")
	err := db.MenuCollection.FindOne(ctx, bson.D{{"menuId", menuId}}).Decode(&menu)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error finding menu"})
	}
	return c.Status(fiber.StatusOK).JSON(menu)
}

func CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error parsing json"})
	}
	if validationError := helpers.ValidateStruct(menu); validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.ID = primitive.NewObjectID()
	menu.MenuId = menu.ID.Hex()

	_, err := db.MenuCollection.InsertOne(ctx, menu)
	if err != nil {
		log.Fatal(err)
	}
	return c.Status(fiber.StatusCreated).JSON(menu)
}

func UpdateMenu(c *fiber.Ctx) error {
	return c.SendString("Update Menu")
}
