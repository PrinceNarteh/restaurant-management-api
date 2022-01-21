package controllers

import (
	"context"
	db "github.com/PrinceNarteh/restaurant-management-api/database"
	"github.com/PrinceNarteh/restaurant-management-api/helpers"
	"github.com/PrinceNarteh/restaurant-management-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)

func GetFoods(c *fiber.Ctx) error {
	return c.SendString("All Foods")
}

func GetFood(c *fiber.Ctx) error {
	var food models.Food
	foodId := c.Params("foodId")

	err := db.FoodCollection.FindOne(ctx, bson.D{{"foodId", foodId}}).Decode(&food)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "food not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": food})
}

func CreateFood(c *fiber.Ctx) error {
	var food models.Food
	var menu models.Menu
	if err := c.BodyParser(&food); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error parsing json"})
	}
	if validationError := helpers.ValidateStruct(food); validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error parsing json"})
	}
	if validationError := helpers.ValidateStruct(menu); validationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validationError})
	}

	if err := db.MenuCollection.FindOne(ctx, bson.D{{"menuId", menu.MenuId}}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "menu not found"})
	}

	food.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Id = primitive.NewObjectID()
	food.FoodId = food.Id.Hex()
	food.MenuId = &menu.MenuId

	foodPrice := helpers.ToFixed(*food.Price, 2)
	food.Price = &foodPrice

	_, err := db.FoodCollection.InsertOne(ctx, food)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error creating food"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": food})
}

func UpdateFood(c *fiber.Ctx) error {
	return c.SendString("Update Food")
}
