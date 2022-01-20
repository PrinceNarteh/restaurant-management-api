package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	db "github.com/PrinceNarteh/restaurant-management-api/database"
	"github.com/PrinceNarteh/restaurant-management-api/helpers"
	"github.com/PrinceNarteh/restaurant-management-api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "enable to parse request"})
	}

	if validateError := helpers.ValidateStruct(user); validateError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": validateError})
	}

	count, err := db.UserCollection.CountDocuments(ctx, bson.D{{"$or", bson.A{bson.D{{"email", user.Email}}, bson.D{{"phoneNumber", user.PhoneNumber}}}}})
	defer cancel()

	if err != nil {
		log.Panic(err)
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User with this email or phone number already exists."})
	}

	user.HashPassword()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()

	accessToken := helpers.GenerateAccessToken(&user)
	refreshToken := helpers.GenerateRefreshToken(&user)
	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken

	_, err = db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error creating user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"accessToken": accessToken, "refreshToken": refreshToken})
}

func Login(c *fiber.Ctx) error {
	var user models.User
	var request struct {
		email    string
		password string
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not pass json"})
	}

	err := db.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "error finding user"})
	}

	isValid := user.ComparePassword(request.password)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid credentails"})
	}

	accessToken := helpers.GenerateAccessToken(&user)
	refreshToken := helpers.GenerateRefreshToken(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"accessToken": accessToken, "refreshToken": refreshToken})
}
