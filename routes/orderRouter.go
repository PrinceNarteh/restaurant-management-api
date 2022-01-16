package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	orderRouter := router.Group("/Orders")
	orderRouter.Get("/", controllers.GetOrders)
	orderRouter.Get("/:orderId", controllers.GetOrder)
	orderRouter.Post("/", controllers.CreateOrder)
	orderRouter.Patch("/:orderId", controllers.UpdateOrder)
}
