package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(router fiber.Router) {
	foodRouter := router.Group("/foods")
	foodRouter.Get("/", controllers.GetFoods)
	foodRouter.Get("/:foodId", controllers.GetFood)
	foodRouter.Post("/", controllers.CreateFood)
	foodRouter.Patch("/:foodId", controllers.UpdateFood)
}
