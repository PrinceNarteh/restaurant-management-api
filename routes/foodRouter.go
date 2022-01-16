package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(router fiber.Router) {
	foodRouter := router.Group("/foods")
	foodRouter.Get("/", controllers.GetAllFoods)
}
