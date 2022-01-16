package routes

import (
	"github.com/PrinceNarteh/restaurant-management-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderItemRoutes(router fiber.Router) {
	menuRouter := router.Group("/orderItems")
	menuRouter.Get("/", controllers.GetOrderItems)
	menuRouter.Get("/:orderItemId", controllers.GetOrderItem)
	menuRouter.Post("/", controllers.CreateOrderItem)
	menuRouter.Patch("/:orderItemId", controllers.UpdateOrderItem)
	router.Get("orderItems-oder/:orderItemId", controllers.GetOrderItemByOrder)
}
