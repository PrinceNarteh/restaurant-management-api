package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	db "github.com/PrinceNarteh/restaurant-management-api/database"
	"github.com/PrinceNarteh/restaurant-management-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

func GetUsers(c *fiber.Ctx) error {

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage
	startIndex, err = strconv.Atoi(c.Query("startIndex"))

	matchStage := bson.D{{"$match", bson.D{{}}}}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
		}}}

	result, err := db.UserCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, projectStage,
	})
	defer cancel()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error occurred while listing user items",
		})
	}

	var allUsers []bson.M
	if err = result.All(ctx, &allUsers); err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(allUsers)
}

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	var user models.User
	err := db.UserCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error occurred while getting user items"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func Register(ctx *fiber.Ctx) error {
	return ctx.SendString("Register User")
}

func Login(ctx *fiber.Ctx) error {
	return ctx.SendString("Login User")
}
